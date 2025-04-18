<script lang="ts">
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'
  import { flip } from 'svelte/animate'
  import CreateServiceQuestion from '../CreateServiceQuestion/CreateServiceQuestion.svelte'
  import CreateServiceQuestionListLayout from './CreateServiceQuestionListLayout.svelte'

  type CreateServiceQuestionListProps = {
    questions: { text: string; id: number }[]
    setQuestions: (value: { text: string; id: number }[]) => void
    removeQuestion: (index: number) => void
  }

  const { questions: items, ...rest }: CreateServiceQuestionListProps = $props()
</script>

{#if items.length === 0}
  <NoContentMessage text="No questions yet" />
{:else}
  <CreateServiceQuestionListLayout {items} {...rest}>
    {#each items as item (item.id)}
      <div class="!outline-none" animate:flip={{ duration: 200 }}>
        <CreateServiceQuestion {...rest} id={item.id} question={item.text} />
      </div>
    {/each}
  </CreateServiceQuestionListLayout>
{/if}
