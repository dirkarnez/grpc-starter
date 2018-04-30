import * as React from 'react';
import { Item, Icon } from 'semantic-ui-react';
import { Lesson } from './proto/mentor_pb';

type StoryListProps = {
  lessons: Lesson.AsObject[],
  selected: Lesson.AsObject | null,
  onLessonSelect: (id: number) => void
};

const StoryList: React.SFC<StoryListProps> = (props) => {
  return (
    <Item.Group divided={true}>
      {props.lessons.map((lesson, i) =>
        <Item
          style={props.selected && lesson.id === props.selected.id
            ? {'backgroundColor': 'rgba(0, 0, 0, 0.08)'}
            : {}
          }
          key={i}
          onClick={() => {
            if (lesson.id) {
              props.onLessonSelect(lesson.id);
            }
          }}
        >
          <Item.Content>
            <Item.Header as="a">{lesson.title}</Item.Header>
            <Item.Extra><Icon  name="star" />{lesson.score} | <Icon  name="user" />{lesson.by}</Item.Extra>
          </Item.Content>
        </Item>
      )}
    </Item.Group>
  );
};

export default StoryList;
