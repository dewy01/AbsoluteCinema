import z from "zod";

export const guestInfoSchema = z.object({
  name: z.string().min(1, 'Imię jest wymagane'),
  email: z.string().min(1, 'Email jest wymagany').email('Nieprawidłowy adres email')
});

export type GuestInfo = z.infer<typeof guestInfoSchema>;