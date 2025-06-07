import { baseUrl } from "@/constants/constants";

export function getResourceUrl(mediaType: "movies" | "reservations", relativePath?: string): string {
  return `${baseUrl}/resources/${mediaType}${relativePath}`;
}
