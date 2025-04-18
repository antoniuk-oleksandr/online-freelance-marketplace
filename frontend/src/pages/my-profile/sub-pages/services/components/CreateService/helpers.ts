import type { StepperItem } from "@/types/StepperItem"

export const makeCreateServiceSteps = (activeStepIndex: number) => {
  const steps: StepperItem[] = [
    {
      text: "Basic Information",
      link: "/my-profile/services/create/basic-information",
    },
    {
      text: "Packages",
      link: "/my-profile/services/create/packages",
    },
    {
      text: "Requirements",
    },
    {
      text: "Media",
    },
  ]

  return steps.map((item, index) => ({
    text: item.text,
    link: index < activeStepIndex ? item.link : undefined,
  }))
}

export const makeCreateServiceTitleData = (activeStepIndex: number) => {
  const titleData = [
    {
      title: "Provide Your Basic Information to Set Up Your Service",
      subtitle: "Enter essential details about your service to get started. This will help potential customers understand what you're offering."
    },
    {
      title: "Define and Customize Your Service Packages",
      subtitle: "Create your own service packages by specifying names, features, and pricing details to best represent your offerings."    
    },
    {
      title: "Define the Specific Requirements for Your Service",
      subtitle: "Outline any prerequisites or criteria that customers must meet in order to avail your service. Be clear about expectations."
    },
    {
      title: "Upload Images to Showcase Your Service",
      subtitle: "Upload high-quality images that represent your service. Visuals help customers get a better understanding of what you provide."
    }
  ];

  return titleData[activeStepIndex] || { title: '', subtitle: '' };
};

