import { z } from 'zod';

const userRegisterSchema = z
  .object({
    email: z.string().email('Nieprawidłowy adres email'),
    password: z.string().min(6, 'Hasło musi mieć co najmniej 6 znaków'),
    confirmPassword: z.string()
  })
  .superRefine((data, ctx) => {
    if (data.password !== data.confirmPassword) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: 'Hasła muszą się zgadzać',
        path: ['confirmPassword']
      });
    }
  });

export type UserRegister = z.infer<typeof userRegisterSchema>;
export { userRegisterSchema };
