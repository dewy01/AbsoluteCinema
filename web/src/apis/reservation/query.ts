import type { components } from '@/types/openapi/reservation';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { useSnackbar } from 'notistack';
import {
  deleteReservationById,
  getReservationById,
  getReservationsByUser,
  postCreateReservation,
  updateReservationPdfPath,
} from './api';

export const useCreateReservation = () => {
  const { enqueueSnackbar } = useSnackbar();
  return useMutation({
    mutationKey: ['createReservation'],
    mutationFn: (data: components['schemas']['CreateReservationInput']) =>
      postCreateReservation(data),
    onSuccess: () => {
      enqueueSnackbar('Rezerwacja została utworzona!');
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Nieprawidłowe dane rezerwacji.', { variant: 'error' });
      } else {
        enqueueSnackbar('Błąd tworzenia rezerwacji.', { variant: 'error' });
      }
    },
  });
};

export const useReservationById = (id: string) => {
  return useQuery({
    queryKey: ['reservation', id],
    queryFn: () => getReservationById(id),
    enabled: !!id,
    retry: false,
  });
};

export const useDeleteReservation = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();
  return useMutation({
    mutationKey: ['deleteReservation', id],
    mutationFn: () => deleteReservationById(id),
    onSuccess: () => {
      enqueueSnackbar('Rezerwacja została usunięta.');
      queryClient.invalidateQueries({ queryKey: ['reservation', id] });
    },
    onError: () => {
      enqueueSnackbar('Nie udało się usunąć rezerwacji.', { variant: 'error' });
    },
  });
};

export const useUpdateReservationPdfPath = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  return useMutation({
    mutationKey: ['updateReservationPdf', id],
    mutationFn: (pdfPath: string) => updateReservationPdfPath(id, pdfPath),
    onSuccess: () => {
      enqueueSnackbar('PDF został zaktualizowany.');
    },
    onError: () => {
      enqueueSnackbar('Nie udało się zaktualizować PDF.', { variant: 'error' });
    },
  });
};

export const useReservationsByUser = (userID: string) => {
  return useQuery({
    queryKey: ['reservationsByUser', userID],
    queryFn: () => getReservationsByUser(userID),
    enabled: !!userID,
    retry: false,
  });
};
