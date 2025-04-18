import type { CreateServicePackage } from "@/types/CreateServicePackage";
import type { CreateServicePackagesTableCell } from "@/types/CreateServicePackagesTableCell";
import { z } from "zod";

const createServicePackageSchema = z.object({
  title: z.string().min(1, 'Title is required.'),
  description: z.string().min(1, 'Description is required.'),
  deliveryDays: z.number().min(1, 'Delivery time must be at least 1 day.'),
  price: z.number().min(1, 'Price must be 0 or more.'),
})

const hasAtLeastOneValidPackage = (pkgs: CreateServicePackage[]) => pkgs.reduce((acc, pkg) =>
  acc = createServicePackageSchema.safeParse(pkg).success, false
)

export const createServicePackagesFormSchema = z.object({
  packages: z
    .array(z.any())
    .refine(
      (pkgs) => pkgs.some(pkg => pkg && createServicePackageSchema.safeParse(pkg).success),
      { message: 'At least one complete package is required.' }  
    ).refine(
      (pkgs) => hasAtLeastOneValidPackage(pkgs),
      { message: 'Incomplete packages must be fully filled out or left empty.' }
    )
})

export const defineCreateServicePackagesTableCells = (): CreateServicePackagesTableCell[][] => [
  [
    { type: 'text' },
    ...['basic', 'standard', 'premium'].map((text, index) => ({
      type: 'text' as const,
      text,
      styles: 'font-semibold uppercase',
      name: `packages.${index}`,
    }))
  ],
  [
    { type: 'text' as const, text: 'Name', styles: 'font-semibold' },
    ...Array.from({ length: 3 }, (_, index) => ({
      type: 'textarea' as const,
      placeholder: 'Name of the package',
      name: `packages.${index}.title`,
    }))
  ],
  [
    { type: 'text' as const, text: 'Description', styles: 'font-semibold' },
    ...Array.from({ length: 3 }, (_, index) => ({
      type: 'textarea' as const,
      placeholder: 'Describe the package',
      name: `packages.${index}.description`,
    }))
  ],
  [
    { type: 'text' as const, text: 'Delivery Time', styles: 'font-semibold' },
    ...Array.from({ length: 3 }, (_, index) => ({
      type: 'input' as const,
      inputType: 'number' as const,
      rightBadge: 'Days',
      name: `packages.${index}.deliveryDays`,
    }))
  ],
  [
    { type: 'text' as const, text: 'Price', styles: 'font-semibold' },
    ...Array.from({ length: 3 }, (_, index) => ({
      type: 'input' as const,
      inputType: 'number' as const,
      rightBadge: '$',
      name: `packages.${index}.price`,
    }))
  ]
]


