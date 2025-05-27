import { errorStore } from "@/common-stores/error-store"
import type { ChatMessage } from "@/types/ChatMessage"
import { ChatMessageType } from "@/types/ChatMessageType"
import type { MyProfileChatRequestResponse } from "@/types/MyProfileChatRequestResponse"
import { z } from "zod"

export const calcChatTextAreaRows = (
  textarea: HTMLTextAreaElement | null,
) => {
  if (!textarea) return

  textarea.style.height = 'auto'
  textarea.rows = 1

  const style = getComputedStyle(textarea)
  const fontSize = parseFloat(style.fontSize)
  let lineHeight = parseFloat(style.lineHeight)

  if (isNaN(lineHeight)) {
    lineHeight = fontSize * 1.2
  }

  const scrollHeight = textarea.scrollHeight
  const newRowCount = Math.ceil(scrollHeight / lineHeight)

  textarea.rows = Math.min(Math.max(newRowCount, 1), 10)
}

export const getUserInformation = () => {
  localStorage.getItem("accessTokenUserId")
  localStorage.getItem("accessTokenUsername")
  localStorage.getItem("accessTokenAvatar")
}

export const makeLastOnlineText = (lastOnline: number) => {
  const lastOnlineDate = new Date(lastOnline);
  const now = new Date();
  const diff = now.getTime() - lastOnlineDate.getTime();

  const seconds = Math.floor(diff / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);
  const months = Math.floor(days / 30);
  const years = Math.floor(months / 12);

  if (years > 0) {
    return `last seen ${years} year${years > 1 ? 's' : ''} ago`;
  } else if (months > 0) {
    return `last seen ${months} month${months > 1 ? 's' : ''} ago`;
  } else if (days > 0) {
    return `last seen ${days} day${days > 1 ? 's' : ''} ago`;
  } else if (hours > 0) {
    return `last seen ${hours} hour${hours > 1 ? 's' : ''} ago`;
  } else if (minutes > 0) {
    return `last seen ${minutes} minute${minutes > 1 ? 's' : ''} ago`;
  } else {
    return `online`;
  }
};

export const formatChatMessageTime = (timestamp: number): string => {
  const date = new Date(timestamp);

  return date.toLocaleTimeString([], {
    hour: '2-digit',
    minute: '2-digit'
  });
};

export const checkIfMessagesOnDifferentDate = (
  prevMessageTimestamp: number,
  newMessageTimestamp: number
): boolean => {
  const oldDate = new Date(prevMessageTimestamp);
  const newDate = new Date(newMessageTimestamp);

  const oldDateYear = oldDate.getFullYear();
  const newDateYear = newDate.getFullYear();
  if (oldDateYear !== newDateYear) return true;

  const oldDateMonth = oldDate.getMonth();
  const newDateMonth = newDate.getMonth();
  if (oldDateMonth !== newDateMonth) return true;

  const oldDateDay = oldDate.getDate();
  const newDateDay = newDate.getDate();
  if (oldDateDay !== newDateDay) return true;

  return false;
}

export const formatDateSeparator = (messageTimestamp: number) => {
  const today = new Date();
  today.setHours(0, 0, 0, 0);

  const messageDate = new Date(messageTimestamp);
  messageDate.setHours(0, 0, 0, 0);

  const timeDiff = today.getTime() - messageDate.getTime();
  const diffDays = Math.floor(timeDiff / (1000 * 3600 * 24));

  if (diffDays === 0) return 'Today';
  if (diffDays === 1) return 'Yesterday';
  if (diffDays < 7) return messageDate.toLocaleDateString('en', { weekday: 'long' });

  return messageDate.toLocaleDateString('en', {
    month: 'short',
    day: 'numeric',
    year: messageDate.getFullYear() !== today.getFullYear() ? 'numeric' : undefined
  });
};

export const checkIfShouldRenderPartnerAvatar = (
  index: number,
  messages: ChatMessage[],
) => {
  if (index === 0) return true;

  const currentMessage = messages[index];
  const prevMessage = messages[index - 1];

  if (currentMessage.senderId !== prevMessage.senderId) return true;
  return checkIfMessagesOnDifferentDate(prevMessage.sentAt, currentMessage.sentAt);
}

export const isFirstMessageOfDay = (
  index: number,
  messages: ChatMessage[]
) => {
  if (index === 0) return true;

  const currentMessage = messages[index];
  const prevMessage = messages[index - 1];

  return checkIfMessagesOnDifferentDate(prevMessage.sentAt, currentMessage.sentAt);
}


export const makeMyProfileChatRequest = async (orderId: string): Promise<MyProfileChatRequestResponse> => {
  const response = {
    data: {
      chatPartner: {
        id: 1,
        username: 'name',
        avatar: 'http://localhost:8030/files/avatar_2.jpg',
        lastOnline: 1811011158000,
      },
      messages: [
        {
          id: 1,
          senderId: 66,
          content: "Hi name, I’d like to order a logo design for my new business. Are you available?",
          sentAt: new Date(Date.UTC(2025, 2, 1, 10, 0, 0)).getTime(),
          files: [],
          type: ChatMessageType.Read,
        },
        {
          id: 2,
          senderId: 1,
          content: "Hi! Yes, I'm available. Could you please share more details about your business and the style you're looking for?",
          sentAt: new Date(Date.UTC(2025, 2, 1, 10, 5, 0)).getTime(),
          files: [],
          type: ChatMessageType.Read,
        },
        {
          id: 3,
          senderId: 66,
          content: "Sure. The business is called 'GreenLeaf Wellness'. I’d like something clean and modern with a natural color palette.",
          sentAt: new Date(Date.UTC(2025, 2, 1, 10, 10, 0)).getTime(),
          files: [],
          type: ChatMessageType.Read,
        },
        {
          id: 4,
          senderId: 1,
          content: "Great! I’ll get started on a few initial concepts. You’ll have the first draft by Friday. Let me know if you have any references or logos you like.",
          sentAt: new Date(Date.UTC(2025, 2, 1, 10, 20, 0)).getTime(),
          files: [],
          type: ChatMessageType.Read,
        },
        {
          id: 5,
          senderId: 66,
          content: "Sounds good. I’ll send over a couple of samples I like. Looking forward to seeing your ideas!",
          sentAt: new Date(Date.UTC(2025, 2, 1, 10, 30, 0)).getTime(),
          files: [],
          type: ChatMessageType.Sent,
        },
      ],
    },
    status: 200,
  } as MyProfileChatRequestResponse;

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error });
  }

  return response;
}


export const chatFormSchema = z.object({
  message: z.string().min(1, { message: "Message is required" }),
})
