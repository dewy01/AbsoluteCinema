import { useScreeningsByCinema } from '@/apis/screening';
import { selectedCinemaAtom } from '@/atoms/cinemaAtom';
import { Box, Button, CircularProgress, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { useAtom } from 'jotai';
import { useState } from 'react';
import { ScreeningList } from './ScreeningList';

export const HomeView = () => {
  const [selectedCinema] = useAtom(selectedCinemaAtom);
  const [selectedDate, setSelectedDate] = useState(dayjs().format('YYYY-MM-DD'));

  const { data, isLoading, isError } = useScreeningsByCinema(
    selectedCinema?.id || '',
    selectedDate
  );

  const dateOptions = Array.from({ length: 7 }).map((_, i) =>
    dayjs().add(i, 'day').format('YYYY-MM-DD')
  );

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
            onClick={() => setSelectedDate(date)}>
            {dayjs(date).format('ddd, DD MMM')}
          </Button>
        ))}
      </Box>

      <ScreeningList screenings={data || []} />
    </Box>
  );
};
