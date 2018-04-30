import { applyMiddleware, combineReducers, createStore } from 'redux';
import lessons, { LessonState } from './reducers/lessons';
import { newGrpcMiddleware } from './middleware/grpc';

interface StoreEnhancerState {
}

export interface RootState extends StoreEnhancerState {
  lessons: LessonState;
}

const rootReducers = combineReducers<RootState>({
  lessons,
});

export default createStore(
  rootReducers,
  applyMiddleware(
    newGrpcMiddleware(),
  )
);
