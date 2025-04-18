import { errorStore } from "@/common-stores/error-store"
import type { ChatMessage } from "@/types/ChatMessage"
import { ChatMessageType } from "@/types/ChatMessageType"
import type { MyProfileChatRequestResponse } from "@/types/MyProfileChatRequestResponse"

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

// export const dividerTime = (timestamp: number): string => {
//   const date = new Date(timestamp);
//   const now = new Date();
//   const todayStart = new Date(now);
//   todayStart.setHours(0, 0, 0, 0);

//   const yesterdayStart = new Date(todayStart);
//   yesterdayStart.setDate(yesterdayStart.getDate() - 1);

//   const timeString = date.toLocaleTimeString([], { 
//     hour: '2-digit', 
//     minute: '2-digit' 
//   });

//   // Within last 24 hours
//   if (date >= todayStart) {
//     return timeString;
//   }

//   // Yesterday
//   if (date >= yesterdayStart) {
//     return `Yesterday, ${timeString}`;
//   }

//   const diffInDays = Math.floor((now.getTime() - date.getTime()) / (1000 * 3600 * 24));
//   const formatter = new Intl.RelativeTimeFormat('en', { numeric: 'auto' });

//   // Last 7 days
//   if (diffInDays <= 7) {
//     return formatter.format(-diffInDays, 'day');
//   }

//   // Last 30 days
//   if (diffInDays <= 30) {
//     const weeks = Math.floor(diffInDays / 7);
//     return formatter.format(-weeks, 'week');
//   }

//   // Within current year
//   if (date.getFullYear() === now.getFullYear()) {
//     return date.toLocaleDateString('en', { 
//       month: 'short', 
//       day: 'numeric' 
//     });
//   }

//   // Older than current year
//   return date.toLocaleDateString('en', { 
//     month: 'short', 
//     day: 'numeric', 
//     year: 'numeric' 
//   });
// };

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
        id: 2,
        firstName: 'John',
        surname: 'Doe',
        avatar: 'http://localhost:8030/files/avatar_2.jpg',
        lastOnline: 1133024292 * 1000,
      },
      messages: [
        {
          id: 1,
          senderId: 66,
          content: "Hey, how's it going?",
          sentAt: 1134641958000,
          files: [],
          type: ChatMessageType.Read,
        },
        {
          id: 2,
          senderId: 2,
          content: "I'm good, thanks! How about you? I'm good, thanks! How about you? I'm good, thanks! How about you? I'm good, thanks! How about you?",
          sentAt: 1711011158000,
          files: [],
          type: ChatMessageType.Read,
        },
        {
          id: 3,
          senderId: 2,
          content: "GtG, see you later!",
          sentAt: 1742740304000,
          files: [],
          type: ChatMessageType.Read,
        },
        {
          id: 4,
          senderId: 66,
          content: "ok!",
          sentAt: 1742826704000,
          files: [],
          type: ChatMessageType.Read,
        },
      ],
    },
    status: 200,
  } as MyProfileChatRequestResponse

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error })
  }

  return response;
}
