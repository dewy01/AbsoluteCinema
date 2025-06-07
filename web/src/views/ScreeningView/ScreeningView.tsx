import { useScreeningById } from '@/apis/screening';
import { getResourceUrl } from '@/utils/resources';
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
    <Box
      p={4}
      display="flex"
      justifyContent="center"
      alignItems="flex-start"
      gap={4}
      flexWrap="wrap"
      maxWidth={900}
      mx="auto">
      <Box
        component="img"
        src={getResourceUrl('movies', data.movie?.photoPath)}
        alt={data.movie?.title}
        sx={{
          width: 300,
          maxWidth: '100%',
          height: 'auto',
          borderRadius: 2,
          backgroundColor: '#eee',
          flexShrink: 0
        }}
      />

      <Box flex="1" minWidth={280} display="flex" flexDirection="column" gap={2}>
        <Typography variant="h3" component="h1" gutterBottom>
          {data.movie?.title || 'Nieznany film'}
        </Typography>

        <Typography variant="subtitle1" color="text.secondary">
          Reżyser: {data.movie?.director || 'Brak danych'}
        </Typography>

        <Typography variant="body1" paragraph>
          {data.movie?.description || 'Brak opisu.'}
        </Typography>

        <Typography variant="body1">
          <strong>Sala:</strong> {data.room?.name}
        </Typography>
        <Typography variant="body1">
          <strong>Start:</strong> {dayjs(data.startTime).format('DD MMM YYYY, HH:mm')}
        </Typography>

        <Box mt={3} display="flex" gap={2} flexWrap="wrap">
          <NavLink to={`/movie/${data.movie?.id}`} style={{ textDecoration: 'none' }}>
            <Button variant="outlined">Przejdź do filmu</Button>
          </NavLink>
          <NavLink to={`/reservation/${data.id}`} style={{ textDecoration: 'none' }}>
            <Button variant="contained" color="primary">
              Rezerwuj
            </Button>
          </NavLink>
        </Box>
      </Box>
    </Box>
  );
};
