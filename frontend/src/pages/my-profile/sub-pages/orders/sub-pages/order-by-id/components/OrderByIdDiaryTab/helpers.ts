import { errorStore } from "@/common-stores/error-store";
import type { DiaryItem } from "@/types/DiaryItem";
import type { MyProfileDiaryRequestResponse } from "@/types/MyProfileDiaryRequestResponse";


export const makeMyProfileDiaryRequest = async (orderId: string): Promise<MyProfileDiaryRequestResponse> => {
  const response = {
    data: {
      dates: [
        {
          date: (new Date(Date.UTC(2025, 2, 1))).getTime(),
          text: "Initial consultation with client. Discussed business goals and logo preferences.",
          bgColor: "ceeec5",  // light green for initial meeting
          textColor: "000000",
        },
        {
          date: (new Date(Date.UTC(2025, 2, 2))).getTime(),
          text: "Client shared logo reference materials and style preferences.",
          bgColor: "fff8c2",  
          textColor: "000000",
        },
        {
          date: (new Date(Date.UTC(2025, 2, 3))).getTime(),
          text: "Created and shared first draft of the logo for feedback.",
          bgColor: "cce5ff",  
          textColor: "000000",
        },
        {
          date: (new Date(Date.UTC(2025, 2, 4))).getTime(),
          text: "Received feedback from client, revision requests made.",
          bgColor: "ffcccb",  
          textColor: "000000",
        }
      ],
      firstDate: (new Date(Date.UTC(2025, 2, 30))).getTime(),
      lastDate: (new Date(Date.UTC(2025, 3, 4))).getTime(),
    },
    status: 200,
  } as MyProfileDiaryRequestResponse;

  if (response.status === 200) {
    response.data.dates = getDiaryMonth(3, 2025, response.data.dates);
  } else {
    errorStore.set({ shown: true, error: response.data.error });
  }
  return response;
}

export const getDiaryMonth = (
  month: number,
  year: number,
  responseDays: DiaryItem[]
): DiaryItem[] => {
  const monthDays = getMonthDays(month, year);
  const daysBefore = getAdditionalDays('before', monthDays[0].date);
  const daysAfter = getAdditionalDays('after', monthDays[monthDays.length - 1].date);
  const emptyDays = [...daysBefore, ...monthDays, ...daysAfter];

  adjustDiaryDays(emptyDays, responseDays);

  return emptyDays;
}

const adjustDiaryDays = (
  emptyDays: DiaryItem[], responseDays: DiaryItem[]
): DiaryItem[] => {
  responseDays.forEach((responseDay) => {
    const index = emptyDays.findIndex((emptyDay) => compareTwoDates(emptyDay.date, responseDay.date));
    if (index !== -1) {
      emptyDays[index] = responseDay;
    }
  });

  return emptyDays;
}

const compareTwoDates = (date1: number, date2: number): boolean => {
  const date1Obj = new Date(date1);
  const date2Obj = new Date(date2);
  const year1 = date1Obj.getUTCFullYear();
  const year2 = date2Obj.getUTCFullYear();
  const month1 = date1Obj.getUTCMonth();
  const month2 = date2Obj.getUTCMonth();
  const day1 = date1Obj.getUTCDate();
  const day2 = date2Obj.getUTCDate();

  return year1 === year2 && month1 === month2 && day1 === day2;
}

const getAdditionalDays = (type: 'before' | 'after', initialDate: number): DiaryItem[] => {
  const items: DiaryItem[] = [];
  const weekDay = new Date(initialDate).getDay();

  for (let i = 0; i < weekDay; i++) {
    const date = new Date(initialDate);
    if (type === 'before') {
      date.setUTCDate(date.getUTCDate() - (weekDay - i));
    } else {
      date.setUTCDate(date.getUTCDate() + (i + 1));
    }
    items.push({
      date: date.getTime(),
    });
  }

  return items;
}

const getMonthDays = (month: number, year: number): DiaryItem[] => {
  const monthDays: DiaryItem[] = [];
  const startDate = new Date(Date.UTC(year, month - 1, 1));
  const endDate = new Date(Date.UTC(year, month + 1, 0));

  let currentDate = startDate;

  while (currentDate <= endDate) {
    monthDays.push({
      date: currentDate.getTime(),
    });

    const nextDate = new Date(currentDate);
    nextDate.setUTCDate(nextDate.getUTCDate() + 1);
    currentDate = nextDate;
  }

  return monthDays;
}
