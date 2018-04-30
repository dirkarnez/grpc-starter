import { RootAction } from '../actions';
import { ADD_LESSON, SELECT_LESSON, LESSONS_INIT } from '../actions/stories';
import { Lesson } from '../proto/mentor_pb';

export type LessonState = {
  readonly lessons: { [lessonId: number]: Lesson.AsObject },
  readonly error: Error | null,
  readonly loading: boolean,
  readonly selected: Lesson.AsObject | null,
};

const initialState = {
  lessons: {},
  error: null,
  loading: false,
  selected: null,
};

export default function (state: LessonState = initialState, action: RootAction): LessonState {

  switch (action.type) {

    case LESSONS_INIT:
      return {...state, loading: true};

    case ADD_LESSON:
      const lesson: Lesson.AsObject = action.payload.toObject();
      const selected = state.selected !== null ? state.selected : lesson;
      if (lesson.id) {
        return {
          ...state,
          loading: false,
          lessons: {...state.lessons, [lesson.id]: lesson},
          selected,
        };
      }
      return state;

    case SELECT_LESSON:
      return {...state, selected: state.lessons[action.payload]};

    default:
      return state;
  }

}
