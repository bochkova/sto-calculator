import { useState, useEffect } from 'react';
import { 
  Box, 
  List, 
  ListItem, 
  ListItemText, 
  Typography,
  CircularProgress,
  ListItemButton
} from '@mui/material';
import axios from 'axios';

interface Calculation {
  id: number;
  name: string;
  formula: string;
}

interface CalculationListProps {
  onSelectCalculation: (calculation: Calculation) => void;
  selectedCalculation: Calculation | null;
}

export default function CalculationList({ onSelectCalculation, selectedCalculation }: CalculationListProps) {
  const [calculations, setCalculations] = useState<Calculation[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCalculations = async () => {
      try {
        const response = await axios.get('/api/calculations');
        setCalculations(response.data);
        setLoading(false);
      } catch (err) {
        setError('Ошибка при загрузке списка расчетов');
        setLoading(false);
        console.error('Error fetching calculations:', err);
      }
    };

    fetchCalculations();
  }, []);

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="200px">
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Typography color="error" sx={{ mt: 2 }}>
        {error}
      </Typography>
    );
  }

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        Доступные расчеты:
      </Typography>
      <List>
        {calculations.map((calculation) => (
          <ListItem key={calculation.id} disablePadding>
            <ListItemButton 
              onClick={() => onSelectCalculation(calculation)}
              selected={selectedCalculation?.id === calculation.id}
            >
              <ListItemText
                primary={calculation.name}
                primaryTypographyProps={{
                  style: { fontWeight: selectedCalculation?.id === calculation.id ? 'bold' : 'normal' }
                }}
              />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Box>
  );
} 