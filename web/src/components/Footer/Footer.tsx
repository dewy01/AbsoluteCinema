import { useDeleteReservation } from '@/apis/reservation';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Link,
  Paper,
  TextField,
  Typography
} from '@mui/material';
import { useState } from 'react';

export const Footer = () => {
  const [open, setOpen] = useState(false);
  const [reservationId, setReservationId] = useState('');
  const { mutateAsync, isPending, isSuccess, isError } = useDeleteReservation(reservationId);

  const handleOpen = () => setOpen(true);
  const handleClose = () => {
    setOpen(false);
    setReservationId('');
  };

  const handleDelete = () => {
    if (!reservationId.trim()) return;

    mutateAsync();
    handleClose();
  };

  return (
    <>
      <Paper
        component="footer"
        sx={{
          p: 2,
          mt: 'auto',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          gap: 1
        }}>
        <Typography variant="subtitle1" fontWeight="bold">
          Rezerwacje
        </Typography>
        <Link component="button" variant="body2" onClick={handleOpen} sx={{ cursor: 'pointer' }}>
          anuluj rezerwacje
        </Link>
      </Paper>

      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Anuluj rezerwację</DialogTitle>
        <DialogContent sx={{ width: '400px' }}>
          <TextField
            label="Numer rezerwacji"
            fullWidth
            variant="outlined"
            value={reservationId}
            onChange={(e) => setReservationId(e.target.value)}
            disabled={isPending}
            autoFocus
            margin="dense"
          />
          {isError && (
            <Typography color="error" mt={1}>
              Nie udało się usunąć rezerwacji. Sprawdź poprawność ID.
            </Typography>
          )}
          {isSuccess && (
            <Typography color="success.main" mt={1}>
              Rezerwacja została anulowana.
            </Typography>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} disabled={isPending}>
            Anuluj
          </Button>
          <Button
            onClick={handleDelete}
            variant="contained"
            color="error"
            disabled={!reservationId.trim() || isPending}>
            Usuń
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
