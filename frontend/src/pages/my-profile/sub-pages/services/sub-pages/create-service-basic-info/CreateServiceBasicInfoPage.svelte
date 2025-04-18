<script lang="ts">
  import CreateService from '../../components/CreateService/CreateService.svelte'
  import CreateServiceBasicInfoFormContent from './components/CreateServiceBasicInfoFormContent/CreateServiceBasicInfoFormContent.svelte'
  import { handleCreateServiceBasicInfoFormSubmit } from './handlers'
  import { createServiceBasicInfoFormSchema } from './helpers'
</script>

<CreateService
  schema={createServiceBasicInfoFormSchema}
  onSubmit={handleCreateServiceBasicInfoFormSubmit}
  activeStepIndex={0}
>
  <CreateServiceBasicInfoFormContent />
</CreateService>
