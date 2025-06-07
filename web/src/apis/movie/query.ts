import type { components } from '@/types/openapi/movie';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { useSnackbar } from 'notistack';
import {
  deleteMovieById,
  getMovieById,
  getMovies,
  postCreateMovie,
  putUpdateMovieById,
} from './api';

export const useMovies = () => {
  return useQuery({
    queryKey: ['movies'],
    queryFn: getMovies,
  });
};

export const useCreateMovie = () => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['createMovie'],
    mutationFn: (data: components['schemas']['CreateMovieInput']) => postCreateMovie(data),
    onSuccess: () => {
      enqueueSnackbar('Movie created successfully!');
      queryClient.invalidateQueries({ queryKey: ['movies'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid movie data.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to create movie.', { variant: 'error' });
      }
    },
  });
};

export const useMovieById = (id: string) => {
  return useQuery({
    queryKey: ['movie', id],
    queryFn: () => getMovieById(id),
  });
};

export const useUpdateMovie = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['updateMovie', id],
    mutationFn: (data: components['schemas']['UpdateMovieInput']) => putUpdateMovieById(id, data),
    onSuccess: () => {
      enqueueSnackbar('Movie updated successfully!');
      queryClient.invalidateQueries({ queryKey: ['movie', id] });
      queryClient.invalidateQueries({ queryKey: ['movies'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid update data.', { variant: 'error' });
      } else if (err instanceof AxiosError && err.response?.status === 404) {
        enqueueSnackbar('Movie not found.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to update movie.', { variant: 'error' });
      }
    },
  });
};

export const useDeleteMovie = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['deleteMovie', id],
    mutationFn: () => deleteMovieById(id),
    onSuccess: () => {
      enqueueSnackbar('Movie deleted successfully!');
      queryClient.invalidateQueries({ queryKey: ['movies'] });
      queryClient.invalidateQueries({ queryKey: ['movie', id] });
    },
    onError: () => {
      enqueueSnackbar('Failed to delete movie.', { variant: 'error' });
    },
  });
};
