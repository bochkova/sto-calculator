import {useState, useEffect} from 'react';
import {
    Box,
    TextField,
    Button,
    Typography,
    Grid,
    Paper,
    CircularProgress,
} from '@mui/material';
import axios from 'axios';
import 'katex/dist/katex.min.css';
import {InlineMath} from 'react-katex';

interface Unit {
    id: number;
    name: string;
    symbol: string;
    type: string;
}

interface Parameter {
    id: number;
    calculation_id: number;
    symbol: string;
    description: string;
    is_constant: boolean;
    value: number | null;
    unit_id: number;
    unit?: Unit;
}

interface CalculationFormProps {
    calculation: {
        id: number;
        name: string;
        formula: string;
    };
}

// Функция для преобразования формулы в формат LaTeX
const convertToLatex = (formula: string): string => {
    let latex = formula;

    // 1. Преобразуем умножение внутри скобок в первую очередь
    latex = latex.replace(/\(([^)]+)\)/g, (match) => {
        let inner = match.slice(1, -1); // убираем внешние скобки
        inner = inner.replace(/\*/g, '\\cdot ');
        return `(${inner})`; // возвращаем скобки
    });

    // 2. Преобразуем дроби вида (Expression)/(Expression) -> \frac{Expression}{Expression}
    // Убедимся, что не заменяем уже существующие \frac
    latex = latex.replace(/\(([^()]+)\)\/\(([^()]+)\)/g, '\\frac{$1}{$2}');

    // 3. Преобразуем простые дроби Term/Term, если они еще остались
    latex = latex.replace(/([A-Za-zА-Яа-я0-9_]+)\/([A-Za-zА-Яа-я0-9_]+)/g, '\\frac{$1}{$2}');

    // 4. Обрабатываем нижние индексы. Ищем букву (латиница или кириллица) за которой следуют
    // один или несколько символов (русские буквы, цифры, '₀').
    // Используем более точное регулярное выражение.
    latex = latex.replace(/([A-Za-zА-Яа-я])([А-Яа-я0-9₀.]+)/g, '\\text{$1}_{\\text{$2}}');


    // 5. Заменяем остальные математические операторы и символы на их LaTeX эквиваленты
    latex = latex
        .replace(/\*/g, '\\cdot ') // умножение (вне скобок и дробей)
        .replace(/\^/g, '^{') // степень (начало)
        .replace(/\s+/g, ' ') // удаляем лишние пробелы
        .replace(/\(/g, '\\left(') // открывающая скобка (для тех, что не в дробях)
        .replace(/\)/g, '\\right)') // закрывающая скобка (для тех, что не в дробях)
        .replace(/([a-zA-ZА-Яа-я]+)_([a-zA-ZА-Яа-я0-9]+)/g, '\\text{$1}_{$2}') // индексы с _ (если есть)


    // 6. Добавляем закрывающие скобки для степеней (если они не закрыты)
    latex = latex.replace(/\^{([^}]*)(?!})}/g, '^{$1}');


    return latex;
};

export default function CalculationForm({calculation}: CalculationFormProps) {
    const [parameters, setParameters] = useState<Parameter[]>([]);
    const [units, setUnits] = useState<Record<number, Unit>>({});
    const [formData, setFormData] = useState<Record<string, string>>({});
    const [result, setResult] = useState<number | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(true);
    const [calculating, setCalculating] = useState(false);

    console.log('Current units state:', units);


    useEffect(() => {
        const fetchParameters = async () => {
            try {
                setLoading(true);
                setResult(null);
                setError(null);
                const parametersResponse = await axios.get(`/api/calculations/${calculation.id}/parameters`);

                console.log('Parameters API response:', parametersResponse.data);

                const unitsMap: Record<number, Unit> = {};
                for (const parameter of parametersResponse.data) {
                    const unitResponse = await axios.get(`/api/unit/${parameter.unit_id}`);
                    unitsMap[unitResponse.data.id] = unitResponse.data;
                }

                console.log('Units map in fetchParameters:', unitsMap);
                setUnits(unitsMap);

                // Добавляем информацию о единицах измерения к параметрам
                const parametersWithUnits = parametersResponse.data.map((param: Parameter) => {
                    const unit = unitsMap[param.unit_id];
                    console.log(`Parameter ${param.symbol} unit_id:`, param.unit_id, 'unit:', unit);
                    return {
                        ...param,
                        unit: unit
                    };
                });

                console.log('Parameters with units:', parametersWithUnits);
                setParameters(parametersWithUnits);

                const initialFormData: Record<string, string> = {};
                parametersWithUnits.forEach((param: Parameter) => {
                    if (!param.is_constant) {
                        initialFormData[param.symbol] = '';
                    }
                });
                setFormData(initialFormData);

                setLoading(false);
            } catch (err) {
                setError('Ошибка при загрузке параметров расчета');
                setLoading(false);
                console.error('Error fetching parameters:', err);
            }
        };

        fetchParameters();
    }, [calculation.id]);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value} = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleCalculate = async () => {
        try {
            setError(null);
            setCalculating(true);
            const paramValues: Record<string, number> = {};

            // Convert form data to numbers and include constant values
            Object.entries(formData).forEach(([name, value]) => {
                if (value === '') {
                    throw new Error(`Поле ${name} не может быть пустым`);
                }
                const numValue = parseFloat(value);
                if (isNaN(numValue)) {
                    throw new Error(`Поле ${name} должно быть числом`);
                }
                paramValues[name] = numValue;
            });

            // Add constant parameters
            parameters.forEach((param: Parameter) => {
                if (param.is_constant && param.value !== null) {
                    paramValues[param.symbol] = param.value;
                }
            });

            const response = await axios.post(`/api/calculations/${calculation.id}/execute`, {
                parameters: paramValues
            });

            if (typeof response.data !== 'number') {
                throw new Error('Неверный формат ответа от сервера');
            }

            setResult(response.data);
            setCalculating(false);
        } catch (err: any) {
            console.error('Calculation error:', err);
            if (axios.isAxiosError(err)) {
                if (err.response?.status === 400) {
                    setError(err.response.data || 'Ошибка в параметрах расчета');
                } else if (err.response?.status === 500) {
                    setError('Ошибка сервера при выполнении расчета');
                } else {
                    setError(err.message || 'Ошибка при выполнении расчета');
                }
            } else {
                setError(err.message || 'Ошибка при выполнении расчета');
            }
            setCalculating(false);
        }
    };

    if (loading) {
        return (
            <Box display="flex" justifyContent="center" alignItems="center" minHeight="200px">
                <CircularProgress/>
            </Box>
        );
    }

    return (
        <Box sx={{
            mt: 3,
            width: '100%',
            maxWidth: '100%',
            px: {xs: 2, sm: 3, md: 4}
        }}>
            <Typography variant="h6" gutterBottom sx={{wordBreak: 'break-word'}}>
                {calculation.name}
            </Typography>
            <Box sx={{
                mb: 3,
                p: 2,
                bgcolor: 'background.paper',
                borderRadius: 1,
                boxShadow: 1
            }}>
                <Typography
                    variant="h6"
                    color="text.secondary"
                    gutterBottom
                    sx={{
                        wordBreak: 'break-word',
                        '& .katex': {
                            fontSize: '1.2em !important'
                        }
                    }}
                >
                    <InlineMath math={convertToLatex(calculation.formula)}/>
                </Typography>
            </Box>

            <Grid
                container
                spacing={{xs: 2, sm: 3}}
                sx={{
                    width: '100%',
                    m: 0
                }}
            >
                {parameters.map((param) => (
                    <Grid
                        item
                        xs={12}
                        sm={6}
                        md={4}
                        key={param.id}
                    >
                        <TextField
                            fullWidth
                            label={<InlineMath math={convertToLatex(param.symbol)}/>}
                            name={param.symbol}
                            type="number"
                            value={param.is_constant ? param.value || '' : (formData[param.symbol] || '')}
                            onChange={handleInputChange}
                            margin="normal"
                            helperText={`${param.description}${param.unit?.symbol ? ` (${param.unit.symbol})` : ''}`}
                            disabled={param.is_constant}
                            InputProps={{
                                readOnly: param.is_constant,
                                endAdornment: param.unit?.symbol && (
                                    <Typography
                                        variant="body2"
                                        color="text.secondary"
                                        sx={{
                                            ml: 1,
                                            display: 'inline',
                                            whiteSpace: 'nowrap'
                                        }}
                                    >
                                        {param.unit.symbol}
                                    </Typography>
                                )
                            }}
                        />
                    </Grid>
                ))}
            </Grid>

            <Box sx={{
                mt: 4,
                mb: 3,
                display: 'flex',
                justifyContent: 'center',
                width: '100%'
            }}>
                <Button
                    variant="contained"
                    color="primary"
                    onClick={handleCalculate}
                    sx={{
                        minWidth: {xs: '100%', sm: '200px'},
                        maxWidth: {sm: '300px'}
                    }}
                    disabled={calculating}
                >
                    {calculating ? 'Вычисление...' : 'Рассчитать'}
                </Button>
            </Box>

            {error && (
                <Typography
                    color="error"
                    sx={{
                        mt: 2,
                        textAlign: 'center',
                        wordBreak: 'break-word'
                    }}
                >
                    {error}
                </Typography>
            )}

            {result !== null && (
                <Paper
                    elevation={2}
                    sx={{
                        p: {xs: 2, sm: 3},
                        mt: 3,
                        mb: 3,
                        textAlign: 'center'
                    }}
                >
                    <Typography variant="h6" gutterBottom>
                        Результат расчета:
                    </Typography>
                    <Typography
                        variant="h4"
                        color="primary"
                        sx={{
                            wordBreak: 'break-word',
                            fontSize: {xs: '1.5rem', sm: '2rem', md: '2.5rem'}
                        }}
                    >
                        {result.toLocaleString('ru-RU', {
                            minimumFractionDigits: 2,
                            maximumFractionDigits: 2
                        })}
                    </Typography>
                </Paper>
            )}
        </Box>
    );
}