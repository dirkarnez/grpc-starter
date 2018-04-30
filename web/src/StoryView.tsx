import * as React from 'react';
import { Lesson } from './proto/mentor_pb';

type StoryViewProps = {
  lesson: Lesson.AsObject,
};

const StoryView: React.SFC<StoryViewProps> = (props) => {
  /*const url = `http://localhost:8900/article-proxy?q=${encodeURIComponent()}`;*/  
  return (
    <h2>{props.lesson.url}</h2>
  );
};

export default StoryView;
