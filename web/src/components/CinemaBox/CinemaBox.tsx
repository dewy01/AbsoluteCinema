import { useCinemas } from '@/apis/cinema';
import { selectedCinemaAtom, writeSelectedCinemaAtom } from '@/atoms/cinemaAtom';
import { Box, Button, CircularProgress, Grid, Paper, Typography } from '@mui/material';
import { useAtom } from 'jotai';

interface Cinema {
  id?: string;
  name?: string;
  address?: string;
  roomIDs?: string[];
}

interface CinemaBoxProps {
  onClose: () => void;
}

export function CinemaBox({ onClose }: CinemaBoxProps) {
  const { data: cinemas, isLoading, isError } = useCinemas();
  const [selectedCinema, setSelectedCinema] = useAtom(selectedCinemaAtom);
  const [, setSelectedCinemaWrite] = useAtom(writeSelectedCinemaAtom);

  const handleSelect = (cinema: Cinema) => {
    setSelectedCinemaWrite(cinema);
    setSelectedCinema(cinema);
    onClose();
  };

  if (isLoading) {
    return (
      <Box
        sx={{
          width: '100px',
          height: '100px',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center'
        }}>
        <CircularProgress />
      </Box>
    );
  }

  if (isError || !cinemas) {
    return null;
  }

  return (
    <Paper
      sx={{
        display: 'flex',
        flexDirection: 'column',
        minWidth: 300,
        px: 4,
        py: 2,
        gap: 2
      }}>
      <Typography variant="h6">Wybierz kino</Typography>

      <Grid container spacing={2}>
        {cinemas.map((cinema) => (
          <Button
            key={cinema.id}
            variant={selectedCinema?.id === cinema.id ? 'contained' : 'outlined'}
            fullWidth
            onClick={() => handleSelect(cinema)}>
            {cinema.name}
          </Button>
        ))}
      </Grid>
    </Paper>
  );
}
