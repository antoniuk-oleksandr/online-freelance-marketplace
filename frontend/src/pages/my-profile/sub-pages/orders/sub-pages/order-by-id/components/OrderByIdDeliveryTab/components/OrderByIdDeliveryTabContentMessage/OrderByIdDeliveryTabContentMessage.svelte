<script lang="ts">
  import OrderByIdDeliveryTabCompleteContentMessageLayout from './OrderByIdDeliveryTabContentMessageLayout.svelte'

  type OrderByIdDeliveryTabContentMessageProps = {
    message: string
    title: string
  }

  const { message, title }: OrderByIdDeliveryTabContentMessageProps = $props()
</script>

<OrderByIdDeliveryTabCompleteContentMessageLayout>
  <h3 class=" text-light-palette-text-primary dark:text-dark-palette-text-primary text-base font-semibold uppercase tracking-wider">{title}</h3>
  <p class=" whitespace-pre-line">{message}</p>
</OrderByIdDeliveryTabCompleteContentMessageLayout>
