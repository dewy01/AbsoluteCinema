import { Box, Button, Typography } from '@mui/material';
import { useNavigate } from 'react-router-dom';

export const NotFoundView = () => {
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
        404
      </Typography>
      <Typography variant="h5" gutterBottom>
        Strona nie została znaleziona
      </Typography>
      <Typography variant="body1" mb={3}>
        Przepraszamy, ale strona, której szukasz, nie istnieje.
      </Typography>
      <Button variant="contained" onClick={() => navigate('/')}>
        Powrót do strony głównej
      </Button>
    </Box>
  );
};
