import { useScreeningById } from '@/apis/screening';
import { Box, Button, CircularProgress, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { NavLink, useParams } from 'react-router-dom';

export const ScreeningView = () => {
  const { id } = useParams<{ id: string }>();
  const { data, isLoading, isError } = useScreeningById(id || '');

  if (isLoading) {
    return (
      <Box p={4} textAlign="center">
        <CircularProgress />
      </Box>
    );
  }

  if (isError || !data) {
    return (
      <Box p={4}>
        <Typography color="error">Nie udało się załadować seansu.</Typography>
      </Box>
    );
  }

  return (
    <Box p={4}>
      <Typography variant="h4" gutterBottom>
        {data.movie?.title || 'Nieznany film'}
      </Typography>
      <Typography variant="subtitle1" gutterBottom>
        Reżyser: {data.movie?.director || 'Brak danych'}
      </Typography>
      <Typography variant="body1" paragraph>
        {data.movie?.description || 'Brak opisu.'}
      </Typography>

      <Box
        component="img"
        src={data.movie?.photoPath || '/placeholder-movie.png'}
        alt={data.movie?.title}
        sx={{
          width: '100%',
          maxWidth: 500,
          height: 'auto',
          borderRadius: 2,
          mb: 3,
          backgroundColor: '#eee'
        }}
      />

      <Typography>
        <strong>Sala:</strong> {data.room?.name}
      </Typography>
      <Typography>
        <strong>Start:</strong> {dayjs(data.startTime).format('DD MMM YYYY, HH:mm')}
      </Typography>

      <Box mt={4} display="flex" gap={2}>
        <NavLink to={`/movie/${data.movie?.id}`}>
          <Button variant="outlined">Przejdź do filmu</Button>
        </NavLink>
        <NavLink to={`/reservation/${data.id}`}>
          <Button variant="contained" color="primary">
            Rezerwuj
          </Button>
        </NavLink>
      </Box>
    </Box>
  );
};
