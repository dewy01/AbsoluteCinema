import { useMovieById } from '@/apis/movie';
import { Box, Chip, CircularProgress, Paper, Typography } from '@mui/material';
import { useParams } from 'react-router-dom';

export const MovieView = () => {
  const { id } = useParams<{ id: string }>();
  const { data: movie, isLoading, isError } = useMovieById(id!);

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" height="300px">
        <CircularProgress />
      </Box>
    );
  }

  if (isError || !movie) {
    return (
      <Box p={4}>
        <Typography color="error">Movie not found.</Typography>
      </Box>
    );
  }

  return (
    <Box p={4}>
      <Typography variant="h4" mb={2}>
        {movie.title}
      </Typography>
      <Typography variant="subtitle1" color="textSecondary" mb={2}>
        Directed by {movie.director}
      </Typography>

      <Paper sx={{ p: 2, mb: 3 }}>
        <Box
          component="img"
          src={movie.photoPath || '/placeholder-movie.png'}
          alt={movie.title}
          sx={{
            width: '100%',
            maxHeight: 400,
            objectFit: 'cover',
            borderRadius: 2
          }}
        />
      </Paper>

      {movie.description && (
        <Typography variant="body1" mb={2}>
          {movie.description}
        </Typography>
      )}

      {movie.actorIDs && movie.actorIDs.length > 0 && (
        <Box mt={2}>
          <Typography variant="subtitle2" gutterBottom>
            Actors:
          </Typography>
          <Box display="flex" gap={1} flexWrap="wrap">
            {movie.actorIDs.map((actorId) => (
              <Chip key={actorId} label={`Actor: ${actorId}`} />
            ))}
          </Box>
        </Box>
      )}
    </Box>
  );
};
