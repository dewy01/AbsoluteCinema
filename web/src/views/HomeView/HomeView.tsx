import { useMovies } from '@/apis/movie';
import { selectedCinemaAtom } from '@/atoms/cinemaAtom';
import { Box, Button, CircularProgress, Grid, Paper, Typography } from '@mui/material';
import { useAtom } from 'jotai';
import { NavLink } from 'react-router-dom';

export const HomeView = () => {
  const [selectedCinema] = useAtom(selectedCinemaAtom);
  const { data: movies, isLoading, isError } = useMovies();

  if (!selectedCinema) {
    return (
      <Box p={4} textAlign="center">
        <Typography variant="h6" color="textSecondary">
          Aby przeglądać filmy, wybierz kino.
        </Typography>
      </Box>
    );
  }

  if (isLoading) {
    return (
      <Box
        sx={{
          width: '100%',
          height: 300,
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center'
        }}>
        <CircularProgress />
      </Box>
    );
  }

  if (isError || !movies) {
    return (
      <Box p={4}>
        <Typography color="error">Brak filmów dla danego kina.</Typography>
      </Box>
    );
  }

  return (
    <Box p={4}>
      <Typography variant="h5" mb={3}>
        Filmy dostępne w kinie: <strong>{selectedCinema.name}</strong>
      </Typography>

      {movies.length === 0 ? (
        <Typography>No movies found for this cinema.</Typography>
      ) : (
        <Grid container spacing={3}>
          {movies.map((movie) => (
            <Paper key={movie.id} elevation={3} sx={{ p: 2 }}>
              <Typography variant="h6" gutterBottom noWrap>
                {movie.title}
              </Typography>
              <Typography variant="subtitle2" color="textSecondary" noWrap>
                Directed by {movie.director}
              </Typography>
              <Box
                component="img"
                src={movie.photoPath || '/placeholder-movie.png'}
                alt={movie.title}
                sx={{
                  width: '100%',
                  height: 180,
                  objectFit: 'cover',
                  mt: 1,
                  borderRadius: 1,
                  backgroundColor: '#eee'
                }}
              />
              <NavLink to={`/movie/${movie.id}`}>
                <Button variant="contained" fullWidth sx={{ mt: 2 }}>
                  Select Movie
                </Button>
              </NavLink>
            </Paper>
          ))}
        </Grid>
      )}
    </Box>
  );
};
