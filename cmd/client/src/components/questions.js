import React, {useState, useEffect } from 'react';
import AnswerCheck from './answerChk';
import NextButton from './nextButton';
import SubmitButton from './submitButton';

export default function Questions(props) {
    const [QuestionList, setQuestions] = useState([]);
    const [NumQuestions, setNumQuestions] = useState(0);
    const [QuestionIdx, setQuestionIdx] = useState(0);

    const [PostResponse, setPostResponse ] = useState([]);
    const [AnswerResp, setAnswerResp] = useState([])

    let questionUrl =  "http://localhost:5000/questions/" + props.urlprefix
    let answerURL = "http://localhost:5000/answer/"
    let title = `<span class="parentCat"> ${props.name}</span> : ${props.subCatName}`

    //Next Question
    function nextQuestion() {
      let newQuestionidx = QuestionIdx + 1
      setQuestionIdx(newQuestionidx)
      setAnswerResp([])
      setPostResponse([])
    }

    //CheckBox Checked
    const handleChange = (e) => {
      const { value, checked } = e.target;
      console.log(`${value} is ${checked}`);

      if (checked) {
        setPostResponse([...PostResponse, value]);
      }

      else {
        setPostResponse(
          PostResponse.filter(
                  (e) => e !== value
              )
          );
      }
    }

    const fetchQuestions = async () => {
        try {
          const response = await fetch(questionUrl, {
            method: 'GET'
          });
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }

          const data = await response.json();
          setQuestions(data);
          let numQuestions = data.length
          setNumQuestions(numQuestions)
        } catch (error) {
          console.error('Error:', error);
        }
      }

    const submitAnswer = async (qid) => {
      let answer = PostResponse.join('');
      console.log("submittedAnswer: " + answer)
      let AnswerURL = `${answerURL}${qid}/${answer}`
      try {
        const response = await fetch(AnswerURL, {
          method: 'GET'
        });
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setAnswerResp([data]);
      } catch (error) {
        console.error('Error:', error);
      }
//      setCorrectAnswerCount((prevValue) => prevValue + 1)
    }

    const displayAnswers = AnswerResp && AnswerResp.map((answer,idx) =>

   <div>
       { answer.IsCorrect &&
        <div key={idx} className='correctAnswer'>

          <p>Correct Answer</p>
        </div>
      }
      { !answer.IsCorrect &&
        <div key={idx} className='incorrectAnswer'>
        <h4>Sorry, that is an <strong>incorrect</strong> Answer.</h4>
        <div>
          <div>
          <h5>Actual Answers</h5>
            {answer.ActualAnswer.map((answer,idx) =>
            <div>
              <p>- {answer.Answer}</p>
            </div>
            )}
        </div>
      </div>
      </div>
      }
      </div>
    );

    const displayQuestion = QuestionList.map((question,idx) =>
    <div>
    <div className="question">
        <p>{question.text}</p>
        <hr ></hr>
        {question.answers.map((answer,idx) =>
        <div>
          <AnswerCheck
            answer={answer}
            idx={idx}
            handleChange={handleChange}
          />
        </div>
        )}
    </div>
    <div className='answer'>
        {displayAnswers}
    </div>
       <SubmitButton
          submitAnswer={submitAnswer}
          qid={question.qid}
       />
      <NextButton
        NextQuestion={nextQuestion}
        NumQuestions={NumQuestions}
        QuestionIdx={QuestionIdx}
      />
       <p className='qid'>Question Id: {question.qid}</p>
    </div>
    )[QuestionIdx];

    useEffect(() => {
      fetchQuestions()
       }, [AnswerResp]);

    return (
    <div>
        <h2 id="list-heading"><span class="parentCat">{props.name} &#92;</span> {props.subCatName}</h2>
        <p>Question: {QuestionIdx + 1} of {NumQuestions}</p>
        <hr></hr>
        {displayQuestion}
      </div>
    );
  }