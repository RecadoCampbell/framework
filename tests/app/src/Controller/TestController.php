<?php
/**
 * Spiral Framework.
 *
 * @license   MIT
 * @author    Anton Titov (Wolfy-J)
 */

namespace Spiral\App\Controller;

use Spiral\Core\Controller;
use Spiral\Router\RouteInterface;
use Spiral\Translator\Traits\TranslatorTrait;

class TestController extends Controller
{
    use TranslatorTrait;

    public function indexAction(string $name = 'Dave')
    {
        return "Hello, {$name}.";
    }

    public function errorAction()
    {
        echo $undefined;
    }

    public function routeAction(RouteInterface $route)
    {
        return $route->getMatches();
    }

    public function requiredAction(int $id)
    {
        //no index
        $this->say(get_class($this));
        $this->say('Hello world');
        $this->say('Hello world', [], 'external');

        l('l');
        l('l', [], 'external');
        p('%s unit|%s units', 10);
        p('%s unit|%s units', 10, [], 'external');

        return $id;
    }
}