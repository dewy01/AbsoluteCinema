import { baseUrl } from '@/constants/constants';
import type { components } from '@/types/openapi/reservation';
import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: baseUrl,
  withCredentials: true,
});

export const postCreateReservation = async (
  data: components['schemas']['CreateReservationInput']
) => {
  const response = await axiosInstance.post('/reservations/', data);
  return response.data as components['schemas']['ReservationOutput'];
};

export const updateReservationById = async (
  id: string,
  data: components['schemas']['UpdateReservationInput']
) => {
  const response = await axiosInstance.put(`/reservations/update/${id}`, data);
  return response.data as components['schemas']['ReservationOutput'];
};

export const getReservationById = async (id: string) => {
  const response = await axiosInstance.get(`/reservations/${id}`);
  return response.data as components['schemas']['ReservationOutput'];
};

export const deleteReservationById = async (id: string) => {
  await axiosInstance.delete(`/reservations/${id}`);
};

export const updateReservationPdfPath = async (
  id: string,
  pdfPath: string
) => {
  await axiosInstance.put(`/reservations/${id}`, { pdfPath });
};

export const getReservationsByUser = async (userID: string) => {
  const response = await axiosInstance.get(`/reservations/user/${userID}`);
  return response.data as components['schemas']['ReservationOutput'][];
};
