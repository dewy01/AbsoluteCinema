import { useSnackbar } from 'notistack';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import type { components } from '@/types/openapi/cinema';
import {
  createCinema,
  deleteCinemaById,
  getAllCinemas,
  getCinemaById,
  updateCinemaById
} from './CinemaApi';

export const useAllCinemas = () => {
  return useQuery({
    queryKey: ['cinemas'],
    queryFn: getAllCinemas
  });
};

export const useCinemaById = (id: string) => {
  return useQuery({
    queryKey: ['cinema', id],
    queryFn: () => getCinemaById(id)
  });
};

export const callCreateCinema = () => {
  const queryClient = useQueryClient();
  const { enqueueSnackbar } = useSnackbar();

  return useMutation({
    mutationKey: ['createCinema'],
    mutationFn: (data: components['schemas']['CreateCinemaInput']) => createCinema(data),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Kino utworzone pomyślnie!' });
      queryClient.invalidateQueries({ queryKey: ['cinemas'] });
    },
    onError: () => {
      enqueueSnackbar({ variant: 'error', message: 'Błąd podczas tworzenia kina.' });
    }
  });
};

export const callUpdateCinema = (id: string) => {
  const queryClient = useQueryClient();
  const { enqueueSnackbar } = useSnackbar();

  return useMutation({
    mutationKey: ['updateCinema', id],
    mutationFn: (data: components['schemas']['UpdateCinemaInput']) => updateCinemaById(id, data),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Kino zaktualizowane!' });
      queryClient.invalidateQueries({ queryKey: ['cinema', id] });
      queryClient.invalidateQueries({ queryKey: ['cinemas'] });
    },
    onError: () => {
      enqueueSnackbar({ variant: 'error', message: 'Błąd podczas aktualizacji kina.' });
    }
  });
};

export const callDeleteCinema = (id: string) => {
  const queryClient = useQueryClient();
  const { enqueueSnackbar } = useSnackbar();

  return useMutation({
    mutationKey: ['deleteCinema', id],
    mutationFn: () => deleteCinemaById(id),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Kino usunięte.' });
      queryClient.invalidateQueries({ queryKey: ['cinemas'] });
      queryClient.invalidateQueries({ queryKey: ['cinema', id] });
    },
    onError: () => {
      enqueueSnackbar({ variant: 'error', message: 'Błąd podczas usuwania kina.' });
    }
  });
};
