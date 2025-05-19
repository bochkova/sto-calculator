declare module 'react-mathjax2' {
  import { ReactNode, ComponentType } from 'react';

  export interface MathJaxContextProps {
    children: ReactNode;
    config?: {
      tex2jax?: {
        inlineMath?: [string, string][];
        displayMath?: [string, string][];
      };
      showProcessingMessages?: boolean;
      messageStyle?: string;
    };
  }

  export interface MathJaxProps {
    children: ReactNode;
    inline?: boolean;
    onRender?: () => void;
  }

  export const MathJaxContext: ComponentType<MathJaxContextProps>;
  export const MathJax: ComponentType<MathJaxProps>;
} 