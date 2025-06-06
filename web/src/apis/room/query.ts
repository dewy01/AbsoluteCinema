import type { components } from '@/types/openapi/room';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { useSnackbar } from 'notistack';
import {
  deleteRoomById,
  getRoomById,
  getRoomsByCinema,
  postCreateRoom,
  putUpdateRoomById,
} from './api';

export const useCreateRoom = () => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['createRoom'],
    mutationFn: (data: components['schemas']['CreateRoomInput']) =>
      postCreateRoom(data),
    onSuccess: () => {
      enqueueSnackbar('Room created successfully!');
      queryClient.invalidateQueries({ queryKey: ['rooms'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid room data.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to create room.', { variant: 'error' });
      }
    },
  });
};

export const useRoomById = (id: string) => {
  return useQuery({
    queryKey: ['room', id],
    queryFn: () => getRoomById(id),
    enabled: !!id,
  });
};

export const useUpdateRoom = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['updateRoom', id],
    mutationFn: (data: components['schemas']['UpdateRoomInput']) =>
      putUpdateRoomById(id, data),
    onSuccess: () => {
      enqueueSnackbar('Room updated successfully!');
      queryClient.invalidateQueries({ queryKey: ['room', id] });
      queryClient.invalidateQueries({ queryKey: ['rooms'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid update data.', { variant: 'error' });
      } else if (err instanceof AxiosError && err.response?.status === 404) {
        enqueueSnackbar('Room not found.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to update room.', { variant: 'error' });
      }
    },
  });
};

export const useDeleteRoom = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['deleteRoom', id],
    mutationFn: () => deleteRoomById(id),
    onSuccess: () => {
      enqueueSnackbar('Room deleted successfully!');
      queryClient.invalidateQueries({ queryKey: ['room', id] });
      queryClient.invalidateQueries({ queryKey: ['rooms'] });
    },
    onError: () => {
      enqueueSnackbar('Failed to delete room.', { variant: 'error' });
    },
  });
};

export const useRoomsByCinema = (cinemaID: string) => {
  return useQuery({
    queryKey: ['roomsByCinema', cinemaID],
    queryFn: () => getRoomsByCinema(cinemaID),
    enabled: !!cinemaID,
  });
};
