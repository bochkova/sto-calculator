import { useState } from 'react';
import { Container, Typography, Box, Drawer, CssBaseline, useTheme } from '@mui/material';
import { styled } from '@mui/material/styles';
import CalculationList from './components/CalculationList';
import CalculationForm from './components/CalculationForm';

interface Calculation {
  id: number;
  name: string;
  formula: string;
}

const drawerWidth = 350;

const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })<{
  open?: boolean;
}>(({ theme, open }) => ({
  flexGrow: 1,
  padding: theme.spacing(3),
  transition: theme.transitions.create('margin', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  [theme.breakpoints.down('sm')]: {
    padding: theme.spacing(2),
  },
  ...(open && {
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginRight: 0,
  }),
}));

const AppBar = styled('div')(({ theme }) => ({
  position: 'fixed',
  right: drawerWidth,
  left: 0,
  top: 0,
  zIndex: theme.zIndex.drawer - 1,
  backgroundColor: theme.palette.background.paper,
  borderBottom: `1px solid ${theme.palette.divider}`,
  padding: theme.spacing(2),
  [theme.breakpoints.down('sm')]: {
    right: 0,
  },
}));

export default function App() {
  const [selectedCalculation, setSelectedCalculation] = useState<Calculation | null>(null);
  const theme = useTheme();

  const handleSelectCalculation = (calculation: Calculation) => {
    setSelectedCalculation(calculation);
  };

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar>
        <Typography variant="h4" component="h1" gutterBottom>
          Калькулятор расчетов
        </Typography>
      </AppBar>
      <Main sx={{ mt: '80px' }}>
        <Container maxWidth="lg" sx={{ px: { xs: 0, sm: 2 } }}>
          {selectedCalculation && (
            <CalculationForm calculation={selectedCalculation} />
          )}
        </Container>
      </Main>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
            position: 'fixed',
            height: '100%',
            top: 0,
            right: 0,
            borderLeft: `1px solid ${theme.palette.divider}`,
            [theme.breakpoints.down('sm')]: {
              width: '100%',
              position: 'fixed',
              height: '100%',
              maxHeight: '100%',
              bottom: 'auto',
              top: 0,
              borderTop: 'none',
              borderLeft: `1px solid ${theme.palette.divider}`,
            },
          },
        }}
        variant="permanent"
        anchor="right"
      >
        <Box sx={{
          p: 2,
          mt: { xs: 0, sm: '80px' },
          height: { xs: '50vh', sm: 'calc(100vh - 80px)' },
          overflow: 'auto'
        }}>
          <CalculationList
            onSelectCalculation={handleSelectCalculation}
            selectedCalculation={selectedCalculation}
          />
        </Box>
      </Drawer>
    </Box>
  );
}