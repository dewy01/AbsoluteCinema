import { CinemaSchedule } from '@/components/CinemaSchedule/CinemaSchedule';
import { useCinema } from '@/contexts';
import { Box } from '@mui/material';

export const HomeView = () => {
  const { selectedCinema } = useCinema();

  return (
    <Box>
      {selectedCinema ? (
        <>
          <CinemaSchedule></CinemaSchedule>
        </>
      ) : (
        <>
          <h1>Brak wybranego kina</h1>
          <p>Proszę wybierz kino, aby kontynuować.</p>
        </>
      )}
    </Box>
  );
};
