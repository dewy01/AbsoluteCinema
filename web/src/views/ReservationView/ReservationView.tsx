import { useCreateReservation } from '@/apis/reservation';
import { useSeatsByScreening } from '@/apis/seat';
import { useAuth } from '@/contexts';
import { zodResolver } from '@hookform/resolvers/zod';
import { Box, Button, CircularProgress, Paper, TextField, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import { useParams } from 'react-router-dom';
import { guestInfoSchema, type GuestInfo } from './schema';

export const ReservationView = () => {
  const { id: screeningID } = useParams<{ id: string }>();
  const { data: seats, isLoading, error } = useSeatsByScreening(screeningID ?? '');
  const { isAuthenticated, userProps: data } = useAuth();
  const createReservationMutation = useCreateReservation();

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isSubmitting }
  } = useForm<GuestInfo>({
    resolver: zodResolver(guestInfoSchema),
    defaultValues: {
      id: data?.id,
      name: data?.name ?? '',
      email: data?.email ?? '',
      seats: []
    }
  });

  const selectedSeats = watch('seats');

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" height="300px">
        <CircularProgress />
      </Box>
    );
  }

  if (error) return <Typography color="error">Nie udało się załadować miejsc.</Typography>;
  if (!seats || seats.length === 0)
    return <Typography>Brak dostępnych miejsc na ten seans.</Typography>;

  const seatsByRow = seats.reduce<Record<string, typeof seats>>((acc, seat) => {
    const row = seat.row || 'Unknown';
    if (!acc[row]) acc[row] = [];
    acc[row].push(seat);
    return acc;
  }, {});
  const sortedRows = Object.keys(seatsByRow).sort();

  const toggleSeat = (seatId: string) => {
    const current = watch('seats');
    if (current.includes(seatId)) {
      setValue(
        'seats',
        current.filter((id) => id !== seatId)
      );
    } else {
      setValue('seats', [...current, seatId]);
    }
  };

  const onSubmit = (guest: GuestInfo) => {
    console.log({
      screeningID: screeningID!,
      reservedSeats: guest.seats.map((seatID) => ({ seatID })),
      userID: guest.id,
      guestName: guest.name,
      guestEmail: guest.email
    });
    createReservationMutation.mutate({
      screeningID: screeningID!,
      reservedSeats: guest.seats.map((seatID) => ({ seatID })),
      userID: guest.id,
      guestName: guest.name,
      guestEmail: guest.email
    });
  };

  return (
    <Box
      p={4}
      display="flex"
      flexDirection={{ xs: 'column', md: 'row' }}
      gap={8}
      alignItems="flex-start"
      justifyContent="center"
      sx={{ overflowX: 'auto' }}>
      <Box flex={1}>
        <Typography variant="h5" gutterBottom>
          Wybierz miejsca
        </Typography>

        {sortedRows.map((row) => (
          <Box key={row} mb={2}>
            <Typography variant="subtitle1" gutterBottom>
              Rząd {row}
            </Typography>

            <Box
              display="inline-flex"
              flexWrap="nowrap"
              gap={1}
              sx={{ whiteSpace: 'nowrap', overflowX: 'auto' }}>
              {seatsByRow[row]
                .sort((a, b) => (a.number ?? 0) - (b.number ?? 0))
                .map((seat) => {
                  const isSelected = selectedSeats.includes(seat.id!);
                  return (
                    <Button
                      key={seat.id}
                      variant={isSelected ? 'contained' : 'outlined'}
                      disabled={seat.isReserved}
                      color={seat.isReserved ? 'error' : isSelected ? 'primary' : 'inherit'}
                      size="small"
                      sx={{ minWidth: 36, minHeight: 36 }}
                      onClick={() => toggleSeat(seat.id!)}>
                      {seat.number}
                    </Button>
                  );
                })}
            </Box>
            {errors.seats && (
              <Typography color="error" mt={1}>
                {errors.seats.message}
              </Typography>
            )}
          </Box>
        ))}
      </Box>

      <Paper
        elevation={4}
        sx={{
          width: '100%',
          minWidth: 350,
          alignSelf: { xs: 'center', md: 'flex-start' },
          p: 4,
          display: 'flex',
          flexDirection: 'column',
          gap: 2
        }}>
        <Typography variant="h5" fontWeight={700} mb={2}>
          Dane rezerwacji
        </Typography>

        <Box
          component="form"
          onSubmit={handleSubmit(onSubmit)}
          noValidate
          autoComplete="off"
          sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
          <TextField
            label="Imię"
            {...register('name')}
            error={!!errors.name}
            helperText={errors.name?.message}
            disabled={isAuthenticated}
            fullWidth
          />

          <TextField
            label="Adres email"
            {...register('email')}
            error={!!errors.email}
            helperText={errors.email?.message}
            disabled={isAuthenticated}
            fullWidth
          />

          <Button
            type="submit"
            variant="contained"
            color="primary"
            disabled={isSubmitting || selectedSeats.length === 0}>
            Zarezerwuj
          </Button>
        </Box>
      </Paper>
    </Box>
  );
};
