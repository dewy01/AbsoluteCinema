import z from "zod";

export const guestInfoSchema = z.object({
  id: z.string().optional(),
  name: z.string().min(1, 'Imię jest wymagane'),
  email: z.string().min(1, 'Email jest wymagany').email('Nieprawidłowy adres email'),
  seats: z.array(z.string()).min(1, 'Musisz wybrać co najmniej jedno miejsce')
});

export type GuestInfo = z.infer<typeof guestInfoSchema>;