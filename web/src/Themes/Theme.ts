import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    mode: 'dark',
    background: {
      default: '#1a1a1a',
      paper: '#222222'
    },
    primary: {
      main: '#ffcc00',
      contrastText: '#000000'
    },
    secondary: {
      main: '#8b0000',
      contrastText: '#ffffff'
    },
    text: {
      primary: '#ffffff',
      secondary: '#bbbbbb'
    }
  },
  typography: {
    fontFamily: '"Montserrat", "Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      fontWeight: 700
    },
    button: {
      textTransform: 'none'
    }
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: '12px'
        }
      }
    }
  }
});

export default theme;
