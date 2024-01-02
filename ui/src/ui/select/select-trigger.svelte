<script lang="ts">
  import { Select as SelectPrimitive } from 'bits-ui';
  import { selectTriggerVariants } from './variants';

  type $$Props = SelectPrimitive.TriggerProps;
  type $$Events = SelectPrimitive.TriggerEvents;

  let className: $$Props['class'] = undefined;
  export { className as class };
</script>

<SelectPrimitive.Trigger
  class={selectTriggerVariants({ className })}
  {...$$restProps}
>
  <slot />
  <div>
    <i class="i-carbon:chevron-sort-down h-4 w-4 opacity-50" />
  </div>
</SelectPrimitive.Trigger>
