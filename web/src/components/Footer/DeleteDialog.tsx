import { useDeleteReservation } from '@/apis/reservation';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
  Typography
} from '@mui/material';
import { useState } from 'react';

interface DeleteDialogProps {
  open: boolean;
  onClose: () => void;
}

export const DeleteDialog = ({ open, onClose }: DeleteDialogProps) => {
  const [reservationId, setReservationId] = useState('');
  const { mutateAsync, isPending, isSuccess, isError } = useDeleteReservation(reservationId);

  const handleDelete = () => {
    if (!reservationId.trim()) return;
    mutateAsync();
    onClose();
    setReservationId('');
  };

  const handleClose = () => {
    onClose();
    setReservationId('');
  };

  return (
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
  );
};
