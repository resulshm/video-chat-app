import { api } from "@/utils";

export async function createRoom(data: FormData) {
  return await api.post("api/v1/room", data);
}
