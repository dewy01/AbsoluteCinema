import type { components } from '@/types/openapi/seat';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { useSnackbar } from 'notistack';
import {
  deleteSeatById,
  getSeatById,
  getSeatsByRoomId,
  getSeatsByScreeningId,
  postCreateSeat,
  putUpdateSeatById,
} from './api';

export const useSeatById = (id: string) => {
  return useQuery({
    queryKey: ['seat', id],
    queryFn: () => getSeatById(id),
    retry: false,
    refetchOnWindowFocus: false,
    enabled: !!id,
  });
};

export const useCreateSeat = () => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['createSeat'],
    mutationFn: (data: components['schemas']['CreateSeatInput']) =>
      postCreateSeat(data),
    onSuccess: () => {
      enqueueSnackbar('Seat created successfully!');
      queryClient.invalidateQueries({ queryKey: ['seats'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid seat data.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to create seat.', { variant: 'error' });
      }
    },
  });
};

export const useUpdateSeat = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['updateSeat', id],
    mutationFn: (data: components['schemas']['UpdateSeatInput']) =>
      putUpdateSeatById(id, data),
    onSuccess: () => {
      enqueueSnackbar('Seat updated successfully!');
      queryClient.invalidateQueries({ queryKey: ['seat', id] });
      queryClient.invalidateQueries({ queryKey: ['seats'] });
    },
    onError: (err) => {
      if (err instanceof AxiosError && err.response?.status === 400) {
        enqueueSnackbar('Invalid update data.', { variant: 'error' });
      } else {
        enqueueSnackbar('Failed to update seat.', { variant: 'error' });
      }
    },
  });
};

export const useDeleteSeat = (id: string) => {
  const { enqueueSnackbar } = useSnackbar();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['deleteSeat', id],
    mutationFn: () => deleteSeatById(id),
    onSuccess: () => {
      enqueueSnackbar('Seat deleted successfully!');
      queryClient.invalidateQueries({ queryKey: ['seats'] });
      queryClient.invalidateQueries({ queryKey: ['seat', id] });
    },
    onError: () => {
      enqueueSnackbar('Failed to delete seat.', { variant: 'error' });
    },
  });
};

export const useSeatsByRoom = (roomID: string) => {
  return useQuery({
    queryKey: ['seatsByRoom', roomID],
    queryFn: () => getSeatsByRoomId(roomID),
    enabled: !!roomID,
  });
};

export const useSeatsByScreening = (screeningID: string) => {
  return useQuery({
    queryKey: ['seatsByScreening', screeningID],
    queryFn: () => getSeatsByScreeningId(screeningID),
    enabled: !!screeningID,
  });
};
