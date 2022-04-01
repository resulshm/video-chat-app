interface Message {
  joined?: Boolean;
  leave?: Boolean;
  type?: "offer" | "answer" | "candidate";
  offer?: RTCSessionDescriptionInit;
  answer?: RTCSessionDescriptionInit;
  candidate?: RTCIceCandidateInit;
}

export type { Message };
