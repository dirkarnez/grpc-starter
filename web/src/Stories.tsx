import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { RootState } from './store';
import { Container, Grid, Header } from 'semantic-ui-react';
import StoryList from './StoryList';
import StoryView from './StoryView';
import { RootAction } from './actions';
import { listLessons, selectLesson } from './actions/stories';
import { Lesson } from './proto/mentor_pb';

type StoriesProps = {
  lessons: Lesson.AsObject[],
  loading: boolean,
  error: Error | null,
  selected: Lesson.AsObject | null,

  fetchLessons: () => void,
  selectLesson: (id: number) => void,
};

class Stories extends React.Component<StoriesProps, {}> {

  componentDidMount() {
    this.props.fetchLessons();
  }

  render() {
    return (
      <Container style={{padding: '1em'}} fluid={true}>
        <Header as="h1" dividing={true}>Hacker News with gRPC-Web</Header>

        <Grid columns={2} stackable={true} divided={'vertically'}>
          <Grid.Column width={4}>
            <StoryList
              selected={this.props.selected}
              lessons={this.props.lessons}
              onLessonSelect={this.props.selectLesson}
            />
          </Grid.Column>

          <Grid.Column width={12} stretched={true}>
            { this.props.selected
              ? <StoryView lesson={this.props.selected} />
              : null
            }
          </Grid.Column>
        </Grid>

      </Container>
    );
  }

}

function mapStateToProps(state: RootState) {
  return {
    lessons: Object.keys(state.lessons.lessons).map(key => state.lessons.lessons[key]),
    loading: state.lessons.loading,
    error: state.lessons.error,
    selected: state.lessons.selected,
  };
}

function mapDispatchToProps(dispatch: Dispatch<RootAction>) {
  return {
    fetchLessons: () => {
      dispatch(listLessons());
    },
    selectLesson: (storyId: number) => {
      dispatch(selectLesson(storyId));
    },
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(Stories);
