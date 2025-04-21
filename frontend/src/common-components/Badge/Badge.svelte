<script lang="ts">
  import BadgeLayout from './BadgeLayout.svelte'
  import BadgeDeleteButton from './components/BadgeDeleteButton/BadgeDeleteButton.svelte'

  type BadgeProps =
    | {
        text: string
        id?: number
        useDeleteButton?: false
      }
    | {
        useDeleteButton: true
        text: string
        deleteFunc: (index: number) => void
        index: number
        isMobile: boolean
        id?: number
      }

  const props: BadgeProps = $props()
</script>

<BadgeLayout id={props.id}>
  <span class="select-none text-base">{props.text}</span>
  {#if props.useDeleteButton}
    <BadgeDeleteButton
      deleteFunc={props.deleteFunc}
      index={props.index}
      isMobile={props.isMobile}
    />
  {/if}
</BadgeLayout>
