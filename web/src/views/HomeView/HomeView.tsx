import { useScreeningsByCinema } from '@/apis/screening';
import { selectedCinemaAtom } from '@/atoms/cinemaAtom';
import { Box, Button, CircularProgress, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { useAtom } from 'jotai';
import { useMemo } from 'react';
import { useSearchParams } from 'react-router-dom';
import { ScreeningList } from './ScreeningList';

export const HomeView = () => {
  const [selectedCinema] = useAtom(selectedCinemaAtom);
  const [searchParams, setSearchParams] = useSearchParams();

  const today = dayjs().format('YYYY-MM-DD');
  const selectedDate = searchParams.get('date') || today;

  const dateOptions = useMemo(
    () => Array.from({ length: 7 }).map((_, i) => dayjs().add(i, 'day').format('YYYY-MM-DD')),
    []
  );

  const { data, isLoading, isError } = useScreeningsByCinema(
    selectedCinema?.id || '',
    selectedDate
  );

  const handleDateChange = (date: string) => {
    setSearchParams({ date });
  };

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

  if (isError) {
    return (
      <Box p={4}>
        <Typography color="error">Błąd podczas ładowania danych.</Typography>
      </Box>
    );
  }

  return (
    <Box p={4}>
      <Typography variant="h5" mb={3}>
        Filmy grane w kinie: <strong>{selectedCinema.name}</strong>
      </Typography>

      <Box display="flex" gap={1} mb={3} flexWrap="wrap">
        {dateOptions.map((date) => (
          <Button
            key={date}
            variant={date === selectedDate ? 'contained' : 'outlined'}
            onClick={() => handleDateChange(date)}
            sx={{ display: 'flex', flexDirection: 'column', gap: 1, p: 1, width: '150px' }}>
            <Typography>{dayjs(date).format('dddd')}</Typography>
            <Typography> {dayjs(date).format('DD MMM')}</Typography>
          </Button>
        ))}
      </Box>

      <ScreeningList screenings={data || []} />
    </Box>
  );
};
