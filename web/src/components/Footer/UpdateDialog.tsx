import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField
} from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

interface UpdateDialogProps {
  open: boolean;
  onClose: () => void;
}

export const UpdateDialog = ({ open, onClose }: UpdateDialogProps) => {
  const [reservationId, setReservationId] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleNavigate = () => {
    if (!reservationId.trim()) {
      setError('Wprowadź poprawny numer rezerwacji.');
      return;
    }
    setError('');
    onClose();
    navigate(`/reservation/${reservationId.trim()}/update`);
    setReservationId('');
  };

  const handleClose = () => {
    onClose();
    setError('');
    setReservationId('');
  };

  return (
    <Dialog open={open} onClose={handleClose}>
      <DialogTitle>Aktualizuj rezerwację</DialogTitle>
      <DialogContent sx={{ width: '400px' }}>
        <TextField
          label="Numer rezerwacji"
          fullWidth
          variant="outlined"
          value={reservationId}
          onChange={(e) => setReservationId(e.target.value)}
          autoFocus
          margin="dense"
          error={!!error}
          helperText={error}
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Anuluj</Button>
        <Button
          onClick={handleNavigate}
          variant="contained"
          color="primary"
          disabled={!reservationId.trim()}>
          Idź do edycji
        </Button>
      </DialogActions>
    </Dialog>
  );
};
