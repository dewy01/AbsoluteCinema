import { useDeleteReservation, useReservationsByUser } from '@/apis/reservation';
import { useCurrentUser } from '@/apis/User';
import { baseUrl } from '@/constants/constants';
import { downloadFileFromPath } from '@/utils/downloadFile';
import { Box, Button, Card, CardContent, CircularProgress, Stack, Typography } from '@mui/material';
import { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { DeleteConfirmationDialog } from './DeleteConfirmationDialog';

export const UserReservationsView = () => {
  const { data: user } = useCurrentUser();
  const { data, isLoading, error } = useReservationsByUser(user?.id ?? '');

  const [selectedIdToDelete, setSelectedIdToDelete] = useState<string | null>(null);

  const { mutate: deleteReservation, isPending } = useDeleteReservation(selectedIdToDelete ?? '');

  const handleDelete = () => {
    if (selectedIdToDelete) {
      deleteReservation();
      setSelectedIdToDelete(null);
    }
  };

  if (isLoading)
    return (
      <Box p={4}>
        <CircularProgress />
      </Box>
    );
  if (error)
    return (
      <Box p={4}>
        <Typography color="error">Nie udało się załadować rezerwacji.</Typography>
      </Box>
    );
  if (!data || data.length === 0)
    return (
      <Box p={4}>
        <Typography>Brak rezerwacji.</Typography>
      </Box>
    );

  return (
    <Box p={4}>
      <Typography variant="h4" gutterBottom>
        Moje rezerwacje
      </Typography>
      <Stack spacing={2}>
        {data.map((reservation) => (
          <Card key={reservation.id}>
            <CardContent>
              <Typography variant="h6">Rezerwacja #{reservation.id}</Typography>
              {reservation.guestName && <Typography>Gość: {reservation.guestName}</Typography>}
              {reservation.guestEmail && <Typography>Email: {reservation.guestEmail}</Typography>}
              <Typography>Miejsca: {reservation.reservedSeats?.length ?? 0}</Typography>
              <Stack direction="row" spacing={2} mt={2}>
                {reservation.pdfPath && (
                  <Button
                    variant="outlined"
                    onClick={() =>
                      downloadFileFromPath(
                        `${baseUrl}/resources/${reservation.pdfPath}`,
                        `rezerwacja-${reservation.id}.pdf`
                      )
                    }>
                    Pobierz PDF
                  </Button>
                )}
                <NavLink to={`/reservation/${reservation.id}/update`}>
                  <Button variant="contained">Edytuj</Button>
                </NavLink>
                <Button
                  variant="outlined"
                  color="error"
                  onClick={() => setSelectedIdToDelete(reservation.id!)}>
                  Usuń
                </Button>
              </Stack>
            </CardContent>
          </Card>
        ))}
      </Stack>

      <DeleteConfirmationDialog
        open={!!selectedIdToDelete}
        isPending={isPending}
        onClose={() => setSelectedIdToDelete(null)}
        onConfirm={handleDelete}
        reservationId={selectedIdToDelete ?? ''}
      />
    </Box>
  );
};
