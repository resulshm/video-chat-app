export interface Message {
  type: "join" | "leave" | "offer" | "answer" | "candidate";
  message: RTCSessionDescriptionInit & RTCIceCandidateInit;
}
