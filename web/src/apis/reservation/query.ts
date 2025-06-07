import { baseUrl } from '@/constants/constants';
import { useAuth } from '@/contexts';
import type { components } from '@/types/openapi/reservation';
import { downloadFileFromPath } from '@/utils/downloadFile';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { useSnackbar } from 'notistack';
import { useNavigate } from 'react-router-dom';
import {
  deleteReservationById,
  getReservationById,
  getReservationsByUser,
  postCreateReservation,
  updateReservationById,
  updateReservationPdfPath,
} from './api';


export const useCreateReservation = () => {
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { isAuthenticated } = useAuth(); 

  return useMutation({
    mutationKey: ['createReservation'],
    mutationFn: (data: components['schemas']['CreateReservationInput']) =>
      postCreateReservation(data),

    onSuccess: async (response) => {
      enqueueSnackbar('Rezerwacja została utworzona!');
      queryClient.invalidateQueries({ queryKey: ['reservation', response.id] });
      queryClient.invalidateQueries({ queryKey: ['reservationsByUser'] });

      const pdfPath = response?.pdfPath;
      if (!pdfPath) {
        console.warn('Brak ścieżki PDF w odpowiedzi serwera.');
        return;
      }
      
      if (!isAuthenticated) {
        try {
          const fullPdfUrl = `${baseUrl}/resources/${pdfPath}`;
          await downloadFileFromPath(fullPdfUrl, `rezerwacja-${response.id}.pdf`).then(()=>navigate('/'));
        } catch (err) {
          enqueueSnackbar('Nie udało się pobrać pliku PDF.', { variant: 'error' });
        }
      } else {
        navigate('/my-reservations');
      }
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

export const useUpdateReservation = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { isAuthenticated } = useAuth();

  return useMutation({
    mutationKey: ['updateReservation', id],
    mutationFn: (data: components['schemas']['UpdateReservationInput']) =>
      updateReservationById(id, data),

    onSuccess: async (response) => {
      enqueueSnackbar('Rezerwacja została zaktualizowana.');
      queryClient.invalidateQueries({ queryKey: ['reservation', id] });
      queryClient.invalidateQueries({ queryKey: ['reservationsByUser'] });

      const pdfPath = response?.pdfPath;
      if (!pdfPath) {
        console.warn('Brak ścieżki PDF w odpowiedzi serwera.');
        return;
      }

      if (!isAuthenticated) {
        try {
          const fullPdfUrl = `${baseUrl}/resources/${pdfPath}`;
          await downloadFileFromPath(fullPdfUrl, `rezerwacja-${response.id}.pdf`).then(() => {
            navigate('/');
          });
        } catch (err) {
          enqueueSnackbar('Nie udało się pobrać pliku PDF.', { variant: 'error' });
        }
      } else {
        navigate('/my-reservations');
      }
    },

    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Nieprawidłowe dane rezerwacji.', { variant: 'error' });
      } else if (err instanceof AxiosError && err.response?.status === 404) {
        enqueueSnackbar('Rezerwacja nie została znaleziona.', { variant: 'error' });
      } else {
        enqueueSnackbar('Błąd podczas aktualizacji rezerwacji.', { variant: 'error' });
      }
    },
  });
};

export const useReservationById = (id: string) => {
  return useQuery({
    queryKey: ['reservation', id],
    queryFn: () => getReservationById(id),
    enabled: !!id,
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
      queryClient.invalidateQueries({ queryKey: ['reservationsByUser'] });
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
  });
};
