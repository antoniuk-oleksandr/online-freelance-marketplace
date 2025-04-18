<script lang="ts">
  type LabeledInputElementLabelProps = {
    label: string
    id: string
  }

  const { label, id }: LabeledInputElementLabelProps = $props()
</script>

<label class="text-light-palette-text-primary dark:text-dark-palette-text-primary" for={id}
  >{label}</label
>
