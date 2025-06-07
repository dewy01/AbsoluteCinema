import { useMovieById } from '@/apis/movie';
import { getResourceUrl } from '@/utils/resources';
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
        <Typography color="error">Nie udało się załadować filmu.</Typography>
      </Box>
    );
  }

  return (
    <Box p={4} maxWidth="md" mx="auto">
      <Typography variant="h3" mb={1}>
        {movie.title}
      </Typography>
      <Typography variant="subtitle1" color="text.secondary" mb={3}>
        Reżyser: {movie.director}
      </Typography>

      <Paper
        elevation={3}
        sx={{
          width: 'fit-content',
          maxWidth: 300,
          mx: 'auto',
          mb: 4,
          borderRadius: 2,
          overflow: 'hidden',
          backgroundColor: '#f9f9f9'
        }}>
        <Box
          component="img"
          src={getResourceUrl('movies', movie.photoPath)}
          alt={movie.title}
          sx={{
            width: '100%',
            height: 'auto',
            display: 'block'
          }}
        />
      </Paper>

      {movie.description && (
        <Typography variant="body1" mb={3}>
          {movie.description}
        </Typography>
      )}

      {movie.actorIDs && movie.actorIDs.length > 0 && (
        <Box mt={2}>
          <Typography variant="subtitle2" gutterBottom>
            Aktorzy:
          </Typography>
          <Box display="flex" gap={1} flexWrap="wrap">
            {movie.actorIDs.map((acotr) => (
              <Chip key={acotr} label={`${acotr}`} />
            ))}
          </Box>
        </Box>
      )}
    </Box>
  );
};
