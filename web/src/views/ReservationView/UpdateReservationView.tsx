import { zodResolver } from '@hookform/resolvers/zod';
import { Box, Button, CircularProgress, Paper, TextField, Typography } from '@mui/material';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { useParams } from 'react-router-dom';

import { useReservationById, useUpdateReservation } from '@/apis/reservation';
import { useSeatsByScreening } from '@/apis/seat';
import { useAuth } from '@/contexts';
import { type GuestInfo, guestInfoSchema } from './schema';

export const UpdateReservationView = () => {
  const { id: reservationID } = useParams<{ id: string }>();
  const { data: reservation, isLoading: loadingReservation } = useReservationById(
    reservationID ?? ''
  );
  const { mutate: updateReservation, isPending: isUpdating } = useUpdateReservation(
    reservationID ?? ''
  );
  const { isAuthenticated, userProps } = useAuth();

  const {
    data: seats,
    isLoading: loadingSeats,
    error: seatsError
  } = useSeatsByScreening(reservation?.screeningID ?? '');

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    reset,
    formState: { errors }
  } = useForm<GuestInfo>({
    resolver: zodResolver(guestInfoSchema),
    defaultValues: {
      id: '',
      name: '',
      email: '',
      seats: []
    }
  });

  useEffect(() => {
    if (reservation) {
      reset({
        id: userProps?.id ?? '',
        name: userProps?.name ?? reservation.guestName ?? '',
        email: userProps?.email ?? reservation.guestEmail ?? '',
        seats: reservation.reservedSeats?.map((seat) => seat.seatID) ?? []
      });
    }
  }, [reservation, userProps, reset]);

  const selectedSeats = watch('seats');

  if (loadingSeats || loadingReservation) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" height="300px">
        <CircularProgress />
      </Box>
    );
  }

  if (seatsError || !seats || !reservation) {
    return <Typography color="error">Nie udało się załadować danych.</Typography>;
  }

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
        current.filter((id) => id !== seatId),
        { shouldValidate: true }
      );
    } else {
      setValue('seats', [...current, seatId], { shouldValidate: true });
    }
  };

  const onSubmit = (guest: GuestInfo) => {
    if (guest.id) {
      updateReservation({
        reservedSeats: guest.seats.map((seatID) => ({ seatID })),
        userID: guest.id,
        guestName: guest.name,
        guestEmail: guest.email
      });
    } else {
      updateReservation({
        reservedSeats: guest.seats.map((seatID) => ({ seatID })),
        guestName: guest.name,
        guestEmail: guest.email
      });
    }
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
          Zaktualizuj miejsca
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
                  const reservedByMe = reservation.reservedSeats?.some(
                    (rs) => rs.seatID === seat.id
                  );
                  const isDisabled = seat.isReserved && !reservedByMe;

                  return (
                    <Button
                      key={seat.id}
                      variant={isSelected ? 'contained' : 'outlined'}
                      disabled={isDisabled}
                      color={isDisabled ? 'error' : isSelected ? 'primary' : 'inherit'}
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
            disabled={isUpdating || selectedSeats.length === 0}>
            Zaktualizuj
          </Button>
        </Box>
      </Paper>
    </Box>
  );
};
