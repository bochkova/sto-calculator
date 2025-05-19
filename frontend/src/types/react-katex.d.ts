declare module 'react-katex' {
  import { ReactNode } from 'react';

  interface MathProps {
    math: string;
    children?: ReactNode;
    block?: boolean;
    errorColor?: string;
    renderError?: (error: Error) => ReactNode;
  }

  export const InlineMath: React.FC<MathProps>;
  export const BlockMath: React.FC<MathProps>;
} 