import type { components } from '@/types/openapi/screening';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { useSnackbar } from 'notistack';
import {
  deleteScreeningById,
  getScreeningById,
  getScreenings,
  getScreeningsByCinema,
  getScreeningsByMovie,
  getScreeningsByRoom,
  postCreateScreening,
  putUpdateScreeningById,
} from './api';

export const useScreenings = (day?: string) => {
  return useQuery({
    queryKey: ['screenings', day],
    queryFn: () => getScreenings(day),
    retry: false,
    refetchOnWindowFocus: false,
  });
};

export const useCreateScreening = () => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['createScreening'],
    mutationFn: (data: components['schemas']['CreateScreeningInput']) =>
      postCreateScreening(data),
    onSuccess: () => {
      enqueueSnackbar('Screening created successfully!');
      queryClient.invalidateQueries({ queryKey: ['screenings'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid screening data.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to create screening.', { variant: 'error' });
      }
    },
  });
};

export const useScreeningById = (id: string) => {
  return useQuery({
    queryKey: ['screening', id],
    queryFn: () => getScreeningById(id),
    retry: false,
    refetchOnWindowFocus: false,
  });
};

export const useUpdateScreening = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['updateScreening', id],
    mutationFn: (data: { startTime: string }) =>
      putUpdateScreeningById(id, data),
    onSuccess: () => {
      enqueueSnackbar('Screening updated successfully!');
      queryClient.invalidateQueries({ queryKey: ['screening', id] });
      queryClient.invalidateQueries({ queryKey: ['screenings'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid update data.', { variant: 'error' });
      } else if (err instanceof AxiosError && err.response?.status === 404) {
        enqueueSnackbar('Screening not found.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to update screening.', { variant: 'error' });
      }
    },
  });
};

export const useDeleteScreening = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['deleteScreening', id],
    mutationFn: () => deleteScreeningById(id),
    onSuccess: () => {
      enqueueSnackbar('Screening deleted successfully!');
      queryClient.invalidateQueries({ queryKey: ['screenings'] });
      queryClient.invalidateQueries({ queryKey: ['screening', id] });
    },
    onError: () => {
      enqueueSnackbar('Failed to delete screening.', { variant: 'error' });
    },
  });
};

export const useScreeningsByMovie = (movieID: string, day?: string) => {
  return useQuery({
    queryKey: ['screeningsByMovie', movieID, day],
    queryFn: () => getScreeningsByMovie(movieID, day),
    enabled: !!movieID,
  });
};

export const useScreeningsByRoom = (roomID: string, day?: string) => {
  return useQuery({
    queryKey: ['screeningsByRoom', roomID, day],
    queryFn: () => getScreeningsByRoom(roomID, day),
    enabled: !!roomID,
  });
};

export const useScreeningsByCinema = (cinemaID: string, day?: string) => {
  return useQuery({
    queryKey: ['screeningsByCinema', cinemaID, day],
    queryFn: () => getScreeningsByCinema(cinemaID, day),
    enabled: !!cinemaID,
  });
};
