import { $api } from "@/plugins";

export async function createRoom(data: FormData) {
  return await $api.post("api/v1/room", data);
}
