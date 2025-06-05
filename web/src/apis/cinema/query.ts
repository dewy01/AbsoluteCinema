import type { components } from '@/types/openapi/cinema';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { useSnackbar } from 'notistack';
import {
  deleteCinemaById,
  getCinemaById,
  getCinemas,
  postCreateCinema,
  putUpdateCinemaById,
} from './api';

export const useCinemas = () => {
  return useQuery({
    queryKey: ['cinemas'],
    queryFn: getCinemas,
    retry: false,
    refetchOnWindowFocus: false,
  });
};

export const useCreateCinema = () => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['createCinema'],
    mutationFn: (data: components['schemas']['CreateCinemaInput']) => postCreateCinema(data),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Cinema created successfully!' });
      queryClient.invalidateQueries({ queryKey: ['cinemas'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar({ variant: 'error', message: 'Invalid cinema data.' });
      } else {
        enqueueSnackbar({ variant: 'error', message: 'Failed to create cinema.' });
      }
    },
  });
};

export const useCinemaById = (id: string) => {
  return useQuery({
    queryKey: ['cinema', id],
    queryFn: () => getCinemaById(id),
    retry: false,
    refetchOnWindowFocus: false,
  });
};

export const useUpdateCinema = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['updateCinema', id],
    mutationFn: (data: components['schemas']['UpdateCinemaInput']) => putUpdateCinemaById(id, data),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Cinema updated successfully!' });
      queryClient.invalidateQueries({ queryKey: ['cinema', id] });
      queryClient.invalidateQueries({ queryKey: ['cinemas'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar({ variant: 'error', message: 'Invalid update data.' });
      } else if (err instanceof AxiosError && err.response?.status === 404) {
        enqueueSnackbar({ variant: 'error', message: 'Cinema not found.' });
      } else {
        enqueueSnackbar({ variant: 'error', message: 'Failed to update cinema.' });
      }
    },
  });
};

export const useDeleteCinema = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['deleteCinema', id],
    mutationFn: () => deleteCinemaById(id),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Cinema deleted successfully!' });
      queryClient.invalidateQueries({ queryKey: ['cinemas'] });
      queryClient.invalidateQueries({ queryKey: ['cinema', id] });
    },
    onError: () => {
      enqueueSnackbar({ variant: 'error', message: 'Failed to delete cinema.' });
    },
  });
};
