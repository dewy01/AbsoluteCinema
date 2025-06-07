import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle
} from '@mui/material';

interface DeleteConfirmationDialogProps {
  open: boolean;
  isPending: boolean;
  onClose: () => void;
  onConfirm: () => void;
  reservationId: string;
}

export const DeleteConfirmationDialog = ({
  open,
  isPending,
  onClose,
  onConfirm,
  reservationId
}: DeleteConfirmationDialogProps) => {
  return (
    <Dialog open={open} onClose={onClose}>
      <DialogTitle>Potwierdzenie usunięcia</DialogTitle>
      <DialogContent>
        <DialogContentText>
          Czy na pewno chcesz usunąć rezerwację o ID: <strong>{reservationId}</strong>?
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Anuluj</Button>
        <Button disabled={isPending} onClick={onConfirm} color="error">
          Usuń
        </Button>
      </DialogActions>
    </Dialog>
  );
};
