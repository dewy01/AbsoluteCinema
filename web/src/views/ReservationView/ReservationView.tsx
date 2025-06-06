import { useCreateReservation } from '@/apis/reservation';
import { useSeatsByScreening } from '@/apis/seat';
import { useCurrentUser } from '@/apis/User';
import { zodResolver } from '@hookform/resolvers/zod';
import { Box, Button, CircularProgress, Grid, Paper, TextField, Typography } from '@mui/material';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { useParams } from 'react-router-dom';
import { guestInfoSchema, type GuestInfo } from './schema';

export const ReservationView = () => {
  const { id: screeningID } = useParams<{ id: string }>();
  const { data: seats, isLoading, error } = useSeatsByScreening(screeningID ?? '');
  const { data } = useCurrentUser();
  const createReservationMutation = useCreateReservation();

  const [selectedSeats, setSelectedSeats] = useState<string[]>([]);

  console.log(data);
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting }
  } = useForm<GuestInfo>({
    resolver: zodResolver(guestInfoSchema),
    defaultValues: {
      name: data?.name ?? '',
      email: data?.email ?? ''
    }
  });

  console.log(seats);

  if (isLoading) return <CircularProgress />;
  if (error) return <Typography color="error">Failed to load seats.</Typography>;
  if (!seats || seats.length === 0)
    return <Typography>No seats found for this screening.</Typography>;

  const seatsByRow = seats.reduce<Record<string, typeof seats>>((acc, seat) => {
    const row = seat.row || 'Unknown';
    if (!acc[row]) acc[row] = [];
    acc[row].push(seat);
    return acc;
  }, {});
  const sortedRows = Object.keys(seatsByRow).sort();

  const toggleSeat = (seatId: string) => {
    if (selectedSeats.includes(seatId)) {
      setSelectedSeats(selectedSeats.filter((id) => id !== seatId));
    } else {
      setSelectedSeats([...selectedSeats, seatId]);
    }
  };

  const onSubmit = (data: GuestInfo) => {
    createReservationMutation.mutate({
      screeningID: screeningID!,
      reservedSeats: selectedSeats.map((seatID) => ({ seatID })),
      guestName: data.name,
      guestEmail: data.email
    });
  };

  return (
    <Box p={4} sx={{ display: 'flex', gap: 4 }}>
      <Box flex={1} maxWidth={400}>
        <Typography variant="h5" gutterBottom>
          Wybierz miejsca
        </Typography>
        {sortedRows.map((row) => (
          <Box key={row} mb={2}>
            <Typography variant="subtitle1" gutterBottom>
              Rząd {row}
            </Typography>
            <Grid container spacing={1}>
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
            </Grid>
          </Box>
        ))}
      </Box>

      <Paper
        elevation={6}
        sx={{
          flex: 1,
          minWidth: 350,
          p: 4,
          display: 'flex',
          flexDirection: 'column',
          gap: 2
        }}>
        <Typography variant="h5" component="h2" fontWeight={700} mb={2}>
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
            disabled={data !== undefined}
            fullWidth
          />

          <TextField
            label="Adres email"
            {...register('email')}
            error={!!errors.email}
            helperText={errors.email?.message}
            disabled={data !== undefined}
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
