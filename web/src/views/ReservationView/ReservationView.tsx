import { Box, Typography } from '@mui/material';
import { useParams } from 'react-router-dom';

export const ReservationView = () => {
  const { id } = useParams<{ id: string }>();

  return (
    <Box p={4}>
      <Typography variant="h5" gutterBottom>
        Rezerwacja seansu
      </Typography>
      <Typography>
        Wybrany seans ID: <strong>{id}</strong>
      </Typography>
    </Box>
  );
};
