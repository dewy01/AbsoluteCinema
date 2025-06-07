import { baseUrl } from "@/constants/constants";

export function getResourceUrl(mediaType: "movies" | "reservations", relativePath?: string): string {
console.log(`${baseUrl}/resources/${mediaType}${relativePath}`)
  return `${baseUrl}/resources/${mediaType}${relativePath}`;
}
