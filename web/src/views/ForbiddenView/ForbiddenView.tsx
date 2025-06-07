import { Box, Button, Typography } from '@mui/material';
import { useNavigate } from 'react-router-dom';

export const ForbiddenView = () => {
  const navigate = useNavigate();

  return (
    <Box
      height="100%"
      display="flex"
      flexDirection="column"
      justifyContent="center"
      alignItems="center"
      textAlign="center"
      p={3}>
      <Typography variant="h2" gutterBottom>
        403
      </Typography>
      <Typography variant="h5" gutterBottom>
        Dostęp zabroniony
      </Typography>
      <Typography variant="body1" mb={3}>
        Nie masz uprawnień, aby zobaczyć tę stronę.
      </Typography>
      <Button variant="contained" onClick={() => navigate('/')}>
        Powrót do strony głównej
      </Button>
    </Box>
  );
};
